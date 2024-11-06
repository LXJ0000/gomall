package code

import (
	"context"
	code "github.com/LXJ0000/gomall/rpc_gen/kitex_gen/code"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Send(ctx context.Context, req *code.SendReq, callOptions ...callopt.Option) (resp *code.SendResp, err error) {
	resp, err = defaultClient.Send(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Send call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Verify(ctx context.Context, req *code.VerifyReq, callOptions ...callopt.Option) (resp *code.VerifyResp, err error) {
	resp, err = defaultClient.Verify(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Verify call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
