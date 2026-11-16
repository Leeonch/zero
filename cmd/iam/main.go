package main

import (
	"flag"
	"net"
	"net/http"

	"example.com/iam-service/internal/config"
	"example.com/iam-service/internal/handler"
	"example.com/iam-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/iam.yaml", "the config file")

func main() {
	flag.Parse()

	var cfg config.Config
	conf.MustLoad(*configFile, &cfg)

	serviceContext := svc.NewServiceContext(cfg)
	httpServer := rest.MustNewServer(cfg.RestConf)
	defer httpServer.Stop()
	httpServer.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/hello",
		Handler: handler.HelloHandler(serviceContext),
	})

	grpcListener, err := net.Listen("tcp", cfg.GRPCListenOn)
	logx.Must(err)
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	proc.AddWrapUpListener(grpcServer.GracefulStop)

	go func() {
		logx.Infof("Starting gRPC server at %s", grpcListener.Addr())
		logx.Must(grpcServer.Serve(grpcListener))
	}()

	logx.Infof("Starting HTTP server at %s:%d", cfg.Host, cfg.Port)
	httpServer.Start()
}
