package main

import (
	"context"
	"log"

	"github.com/go-kratos/etcd/registry"
	pb "github.com/go-kratos/examples/helloworld/helloworld"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	conf := clientv3.Config{}
	conf.Endpoints = []string{"127.0.0.1:2379"}
	cli, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}
	r := registry.New(cli)
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("discovery:///helloworld"),
		transgrpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}
