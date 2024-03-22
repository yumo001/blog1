package main

import (
	"github.com/yumo001/blog1/grpcBegin"
	"github.com/yumo001/blog1/initialize"
	"github.com/yumo001/blog1/logic"
	"google.golang.org/grpc"
)

func init() {
	initialize.Viper()
	initialize.Nacos()
	initialize.Consul()
}

func main() {
	grpcBegin.NewBolgServer(8081, func(s *grpc.Server) {
		logic.RegisterUsersServer(s)
	})
}
