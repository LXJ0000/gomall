package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LXJ0000/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/LXJ0000/gomall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli, err := echo.NewClient("demo_thrift_service",
		client.WithHostPorts("localhost:8888"),                  // 指定客户端要连接的服务器地址和端口
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler), // 处理传输层的元数据头信息
		client.WithTransportProtocol(transport.TTHeader),        // 指定了客户端使用的传输协议
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ // 设置了客户端的基本信息，包括服务名称
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Echo(context.Background(), &api.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
