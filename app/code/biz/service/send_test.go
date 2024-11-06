package service

import (
	"context"
	"testing"
	code "github.com/LXJ0000/gomall/app/code/kitex_gen/code"
)

func TestSend_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendService(ctx)
	// init req and assert value

	req := &code.SendReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
