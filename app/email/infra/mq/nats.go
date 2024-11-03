package mq

import (
	"github.com/LXJ0000/gomall/app/email/utils"
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(nats.DefaultURL)
	utils.MustHandleError(err)
}
