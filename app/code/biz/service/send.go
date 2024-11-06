package service

import (
	"context"
	code "github.com/LXJ0000/gomall/app/code/kitex_gen/code"
)

type SendService struct {
	ctx context.Context
} // NewSendService new SendService
func NewSendService(ctx context.Context) *SendService {
	return &SendService{ctx: ctx}
}

// Run create note info
func (s *SendService) Run(req *code.SendReq) (resp *code.SendResp, err error) {
	// Finish your business logic.

	return
}
