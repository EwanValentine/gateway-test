package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/ewanvalentine/gateway-test/api/proto/greeter"
	proto "github.com/ewanvalentine/gateway-test/proto/greeter"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	log.Println(req.Name)
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterGreeterHandlerFromEndpoint(ctx, mux, ":9090", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe("localhost:8080", mux)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
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
