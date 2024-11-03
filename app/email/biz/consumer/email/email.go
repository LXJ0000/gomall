package email

import (
	"github.com/LXJ0000/gomall/app/email/infra/mq"
	"github.com/LXJ0000/gomall/app/email/infra/notify"
	"github.com/LXJ0000/gomall/app/email/kitex_gen/email"
	"github.com/LXJ0000/gomall/app/email/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}
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
