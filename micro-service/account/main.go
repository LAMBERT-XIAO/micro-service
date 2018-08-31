package main

import (
    "lambert.com.cn/micro-service/core/server"
)

func main() {
    rpcServer := server.NewRPCServer()
    rpcServer.Start()
}
