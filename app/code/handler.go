package main

import (
	"context"
	code "github.com/LXJ0000/gomall/app/code/kitex_gen/code"
	"github.com/LXJ0000/gomall/app/code/biz/service"
)

// CodeServiceImpl implements the last service interface defined in the IDL.
type CodeServiceImpl struct{}

// Send implements the CodeServiceImpl interface.
func (s *CodeServiceImpl) Send(ctx context.Context, req *code.SendReq) (resp *code.SendResp, err error) {
	resp, err = service.NewSendService(ctx).Run(req)

	return resp, err
}

// Verify implements the CodeServiceImpl interface.
func (s *CodeServiceImpl) Verify(ctx context.Context, req *code.VerifyReq) (resp *code.VerifyResp, err error) {
	resp, err = service.NewVerifyService(ctx).Run(req)

	return resp, err
}
