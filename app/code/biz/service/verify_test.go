package service

import (
	"context"
	"testing"
	code "github.com/LXJ0000/gomall/app/code/kitex_gen/code"
)

func TestVerify_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyService(ctx)
	// init req and assert value

	req := &code.VerifyReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
