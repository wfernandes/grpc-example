package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/wfernandes/grpc-example/definitions"

	"google.golang.org/grpc"
)

func main() {
	address := "localhost:8081"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := definitions.NewCounterClient(conn)

	// Contact the server and print out its response.
	name := "defaultName"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	countStream, err := c.Count(context.Background(), &definitions.CountRequest{Name: name})
	if err != nil {
		log.Fatalf("could not Count: %v", err)
	}

	for {
		resp, err := countStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error receiving: %s", err)
		}
		log.Printf("CountResponse: %s\n", resp)
	}
}
