package rpc

import (
	"sync"

	"github.com/LXJ0000/gomall/app/gateway/conf"
	"github.com/LXJ0000/gomall/app/gateway/utils"
	"github.com/LXJ0000/gomall/common/clientsuite"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	c, err := userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: conf.GetConf().Hertz.Address, RegistryAddr: conf.GetConf().Hertz.RegistryAddr}))
	utils.MustHandleError(err)

	UserClient = c
}
