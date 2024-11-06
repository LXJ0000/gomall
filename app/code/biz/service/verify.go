package service

import (
	"context"
	code "github.com/LXJ0000/gomall/app/code/kitex_gen/code"
)

type VerifyService struct {
	ctx context.Context
} // NewVerifyService new VerifyService
func NewVerifyService(ctx context.Context) *VerifyService {
	return &VerifyService{ctx: ctx}
}

// Run create note info
func (s *VerifyService) Run(req *code.VerifyReq) (resp *code.VerifyResp, err error) {
	// Finish your business logic.

	return
}
