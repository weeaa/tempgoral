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
	if amount < MIN_TIP_AMOUNT {
		amount = MIN_TIP_AMOUNT
	}
	return system.NewTransferInstruction(amount, from, to)
}

func (c *Client) GenerateTipInstructionWithRandomTipAddress(amount uint64, from solana.PublicKey) *system.Transfer {
	if amount < MIN_TIP_AMOUNT {
		amount = MIN_TIP_AMOUNT
	}
	return system.NewTransferInstruction(amount, from, solana.MustPublicKeyFromBase58(PickRandomNozomiTipAddress()))
}
