package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/LXJ0000/gomall/demo/demo_proto/conf"
	"github.com/LXJ0000/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/LXJ0000/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/LXJ0000/gomall/demo/demo_proto/middleware"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}
	cli, err := echo.NewClient("demo_proto",
		client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client") // key 必须是大写加下划线风格
	var bizErr *kerrors.GRPCBizStatusError
	resp, err := cli.Echo(ctx, &pbapi.Request{Message: "hello"})
	if err != nil {
		if ok := errors.As(err, &bizErr); ok {
			fmt.Println(bizErr)
			return
		}
		log.Fatal(err)
	}
	log.Println(resp)
}
