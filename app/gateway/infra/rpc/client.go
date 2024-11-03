package rpc

import (
	"sync"

	"github.com/LXJ0000/gomall/app/gateway/conf"
	"github.com/LXJ0000/gomall/app/gateway/utils"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
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
	resolver, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)

	c, err := userservice.NewClient("user", client.WithResolver(resolver))
	utils.MustHandleError(err)

	UserClient = c
}
