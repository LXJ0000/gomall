package code

import (
	"context"
	code "github.com/LXJ0000/gomall/rpc_gen/kitex_gen/code"

	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/code/codeservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() codeservice.Client
	Service() string
	Send(ctx context.Context, Req *code.SendReq, callOptions ...callopt.Option) (r *code.SendResp, err error)
	Verify(ctx context.Context, Req *code.VerifyReq, callOptions ...callopt.Option) (r *code.VerifyResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := codeservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient codeservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() codeservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Send(ctx context.Context, Req *code.SendReq, callOptions ...callopt.Option) (r *code.SendResp, err error) {
	return c.kitexClient.Send(ctx, Req, callOptions...)
}

func (c *clientImpl) Verify(ctx context.Context, Req *code.VerifyReq, callOptions ...callopt.Option) (r *code.VerifyResp, err error) {
	return c.kitexClient.Verify(ctx, Req, callOptions...)
}
