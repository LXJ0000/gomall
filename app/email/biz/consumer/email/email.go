package email

import (
	"context"

	"github.com/LXJ0000/gomall/app/email/infra/mq"
	"github.com/LXJ0000/gomall/app/email/infra/notify"
	"github.com/LXJ0000/gomall/app/email/kitex_gen/email"
	"github.com/LXJ0000/gomall/app/email/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	tracer := otel.Tracer("gomall-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}
		
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		_, span := tracer.Start(ctx, "gomall-email-consumer")
		defer span.End()

		noopEmail := notify.NewNooEmail()
		_ = noopEmail.Send(&req)
	})
	utils.MustHandleError(err)

	// cancel the subscription
	server.RegisterShutdownHook(func() {
		_ = sub.Unsubscribe()
		mq.Nc.Close()
	})
}
