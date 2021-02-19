package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC port")
)

func main() {
	flag.Parse()
	log.Printf("[main] Entered (port: %d", *port)

	endpoint := fmt.Sprintf(":%d", *port)
	server, err := daprd.NewService(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.AddServiceInvocationHandler("echo", echo); err != nil {
		log.Fatal(err)
	}

	log.Printf("[main] Start gRPC service: %s", endpoint)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
func echo(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	log.Println("[main:echo] Entered")

	if in == nil {
		err = errors.New("nil invocation parameter")
		return
	}
	log.Printf(
		"[main:echo] ContentType:%s, Verb:%s, QueryString:%s, DataTypeURL:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.DataTypeURL, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
