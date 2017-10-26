package main

import (
	"fmt"
	"log"
	"time"

	proto "github.com/ewanvalentine/gateway-test/proto/greeter"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	log.Println(req.Name)
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {

	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
