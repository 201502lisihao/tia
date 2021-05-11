package server

import (
	"context"

	"my-project/api"
)

type ApiServer struct {
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (m *ApiServer) Hello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{
		Code: 0,
		Msg:  "hello",
		Data: nil,
	}, nil
}
