package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"

	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"github.com/zeromicro/zero-examples/discovery/consul/server/internal/config"
	"github.com/zeromicro/zero-examples/rpc/remote/unary"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/consul.yaml", "the config file")

type GreetServer struct {
	lock     sync.Mutex
	alive    bool
	downTime time.Time
}

func NewGreetServer() *GreetServer {
	return &GreetServer{
		alive: true,
	}
}

func (gs *GreetServer) Greet(ctx context.Context, req *unary.Request) (*unary.Response, error) {
	fmt.Println("=>", req)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &unary.Response{
		Greet: "hello from " + hostname,
	}, nil
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		unary.RegisterGreeterServer(grpcServer, NewGreetServer())
	})
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		resp, err = handler(ctx, req)
		log.Printf("method: %s time: %v\n", info.FullMethod, time.Since(st))
		return resp, err
	}

	server.AddUnaryInterceptors(interceptor)

	_ = consul.RegitserService(c.ListenOn, c.Consul)

	server.Start()
}
