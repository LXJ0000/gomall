package service

import (
	"context"
	"testing"
	email "github.com/LXJ0000/gomall/app/email/kitex_gen/email"
)

func TestSend_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendService(ctx)
	// init req and assert value

	req := &email.EmailReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
