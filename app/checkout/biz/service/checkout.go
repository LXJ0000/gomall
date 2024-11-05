package service

import (
	"context"

	"github.com/LXJ0000/gomall/app/checkout/infra/mq"
	checkout "github.com/LXJ0000/gomall/app/checkout/kitex_gen/checkout"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/email"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	// TODO ....

	// producer
	data, _ := proto.Marshal(&email.EmailReq{
		To:          "xxx",
		From:        "xxx",
		Subject:     "xxx",
		ContentType: "xxx",
		Body:        "xxx",
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}
	
	// inject trace context
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	return
}
