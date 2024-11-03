package rpc

import (
	"sync"

	"github.com/LXJ0000/gomall/app/user/conf"
	"github.com/LXJ0000/gomall/app/user/utils"
	"github.com/LXJ0000/gomall/common/clientsuite"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient userservice.Client
	once       sync.Once
	err        error
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}
	UserClient, err = userservice.NewClient("user", opts...)
	utils.MustHandleError(err)
}
