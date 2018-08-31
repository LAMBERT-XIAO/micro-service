package server

import (
    "fmt"
    "net"
    "google.golang.org/grpc"
)

type RPCServer struct {
    ServiceName string
    ServiceHost string
    ServicePort int
    Env string
}

var (
    serviceHost = "127.0.0.1"
    servicePort = 1701
    env = "local"
)

func NewRPCServer(serviceName string) *RPCServer {
    return &RPCServer{
        serviceName,
        serviceHost,
        servicePort,
        env,
    }
}

func (server *RPCServer) Start(registerServerFunc interface{}, pService interface{}) {
    listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", servicePort))

    if err != nil {
        panic(err)
    }

    grpc.NewServer(grpc.Un)
}
