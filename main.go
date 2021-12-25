package main

import (
	"gin-gateway/router"
	"gin-gateway/rpc"
)

func main() {
	rpc.Init()
	router.Init()
}
