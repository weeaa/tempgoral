package temporal

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
)

type Client struct {
	Ctx    context.Context
	Client *rpc.Client
	Ste    *rpc.Client
	ApiKey string
}

func New(ctx context.Context, dialURL Region, ste SendTransactionEndpoint, apiKey string) *Client {
	client := Client{
		Ctx:    ctx,
		ApiKey: apiKey,
	}

	client.Client = rpc.New(string(dialURL) + client.ApiKey)
	client.Ste = rpc.New(string(ste) + client.ApiKey)

	return &client
}

func (c *Client) SendTransaction(tx *solana.Transaction) (solana.Signature, error) {
	return c.Ste.SendTransaction(c.Ctx, tx)
}

func (c *Client) GenerateTipInstruction(amount uint64, from solana.PublicKey) *system.Transfer {
	return system.NewTransferInstruction(amount, from, solana.MustPublicKeyFromBase58(PickRandomNozomiTipAddress()))
}
