package temporal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Ctx    context.Context
	Client *rpc.Client
	Ws     *websocket.Conn
	ApiKey string
}

func New(ctx context.Context, region Region, apiKey string) *Client {
	client := Client{
		Ctx:    ctx,
		ApiKey: apiKey,
	}

	client.Client = rpc.New(string(region) + client.ApiKey)

	return &client
}

func (c *Client) SendTransaction(tx *solana.Transaction) (solana.Signature, error) {
	return c.Client.SendTransaction(c.Ctx, tx)
}

func (c *Client) GenerateTipInstruction(amount uint64, from, to solana.PublicKey) *system.Transfer {
	return system.NewTransferInstruction(amount, from, to)
}

func (c *Client) GenerateTipInstructionWithRandomTipAddress(amount uint64, from solana.PublicKey) *system.Transfer {
	return system.NewTransferInstruction(amount, from, solana.MustPublicKeyFromBase58(PickRandomNozomiTipAddress()))
}

type TipInfo struct {
	Time                     string `json:"time"`
	LandedTips25ThPercentile string `json:"landed_tips_25th_percentile"`
	LandedTips50ThPercentile string `json:"landed_tips_50th_percentile"`
	LandedTips75ThPercentile string `json:"landed_tips_75th_percentile"`
	LandedTips95ThPercentile string `json:"landed_tips_95th_percentile"`
	LandedTips99ThPercentile string `json:"landed_tips_99th_percentile"`
}

func GetTipInfo(client *http.Client) (*TipInfo, error) {
	if client == nil {
		client = &http.Client{}
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   "api.nozomi.temporal.xyz",
			Path:   "/tip_floor",
		},
		Header: http.Header{
			"User-Agent": {"tempgoral"},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error getting tip info: got %s response status [%s]", resp.Status, string(body))
	}

	var r TipInfo
	err = json.Unmarshal(body, &r)
	return &r, err
}

func SubscribeTipStream(ctx context.Context) (<-chan []*TipInfo, <-chan error, error) {
	dialer := websocket.Dialer{}
	ch := make(chan []*TipInfo)
	chErr := make(chan error)

	conn, _, err := dialer.Dial(tipStreamURL, nil)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		defer close(ch)
		defer close(chErr)
		defer conn.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				var r []*TipInfo
				if err = conn.ReadJSON(&r); err != nil {
					chErr <- err
					time.Sleep(500 * time.Millisecond)
					continue
				}

				ch <- r
			}
		}
	}()

	return ch, chErr, nil
}
