package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/LXJ0000/gomall/app/user/biz/dal"
	"github.com/LXJ0000/gomall/app/user/conf"
	"github.com/LXJ0000/gomall/app/user/infra/mq"
	"github.com/LXJ0000/gomall/common/mtl"
	"github.com/LXJ0000/gomall/common/serversuite"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	serviceName  = conf.GetConf().Kitex.Service
	registryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// tracing init
	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background()) // 服务关闭之前把剩余的链路都上传完毕 （定时分批上次）
	// mtl init
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, registryAddr)
	// mq init
	mq.Init()
	// init dal
	dal.Init()

	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		fmt.Println(err)
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		}),
	)

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		_ = asyncWriter.Sync()
	})
	return
}
