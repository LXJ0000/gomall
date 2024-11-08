// Code generated by hertz generator.

package main

import (
	"context"
	"time"

	"github.com/LXJ0000/gomall/app/gateway/biz/router"
	"github.com/LXJ0000/gomall/app/gateway/conf"
	"github.com/LXJ0000/gomall/app/gateway/infra/rpc"
	"github.com/LXJ0000/gomall/common/mtl"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	serviceName  = conf.GetConf().Hertz.Service
	registryAddr = conf.GetConf().Hertz.RegistryAddr
)

func main() {
	// load env
	// godotenv.Load()
	// mtl init
	consul, registryInfo := mtl.InitMetric(serviceName, conf.GetConf().Hertz.MetricsPort, registryAddr)
	defer consul.Deregister(registryInfo)
	// tracing init
	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background()) // 服务关闭之前把剩余的链路都上传完毕 （定时分批上次）
	// init rpc client
	rpc.Init()

	tracer, cfg := hertztracing.NewServerTracer()

	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address), server.WithTracer( // 添加 prometheus 中间件到 hertz
		hertzprom.NewServerTracer(
			"",
			"",
			hertzprom.WithDisableServer(true),
			hertzprom.WithRegistry(mtl.Registry),
		),
	),
		tracer,
	)

	h.Use(hertztracing.ServerMiddleware(cfg))

	registerMiddleware(h)

	// add a ping route to test
	// h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	// 	ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	// })
	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}
