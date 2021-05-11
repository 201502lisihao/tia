package app

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"

	"my-project/api"
	"my-project/internal/app/server"
	"my-project/internal/config"
)

type ServerRunner struct {
	conf *config.Config
	apiServer *server.ApiServer
}

func NewServerRunner(
	conf *config.Config,
	apiServer *server.ApiServer,
) *ServerRunner {
	return &ServerRunner{
		conf: conf,
		apiServer: apiServer,
	}
}

func (e *ServerRunner) RunGrpcServer() error {
	s := grpc.NewServer()
	api.RegisterApiServer(s, e.apiServer)

	lis, err := net.Listen("tcp", e.conf.App.Addr.Grpc)
	if err != nil {
		return err
	}
	lis = netutil.LimitListener(lis, 1024)

	log.Println("启动 GRPC 服务", "port", e.conf.App.Addr.Grpc)
	return s.Serve(lis)
}

func (e *ServerRunner) RunHttpServer() error {
	ctx := context.Background()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterApiHandlerFromEndpoint(ctx, mux, e.conf.App.Addr.Grpc, opts)
	if err != nil {
		return err
	}
	log.Println("启动 HTTP 服务", "port", e.conf.App.Addr.Http)
	return http.ListenAndServe(e.conf.App.Addr.Http, mux)
}
