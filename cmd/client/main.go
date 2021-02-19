package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

var (
	appID = flag.String("--app", "", "dapr appID may include namespace")
)

func main() {
	log.Println("[main] Entered")

	client, err := dapr.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        []byte("Hello Freddie"),
	}

	ctx := context.Background()
	for {
		resp, err := client.InvokeMethodWithContent(ctx, *appID, "echo", "post", content)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[main] Response: %s\n", string(resp))
		time.Sleep(5 * time.Second)
	}
}
