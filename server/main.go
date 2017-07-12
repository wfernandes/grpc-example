package main

import (
	"log"
	"net"
	"time"

	"github.com/wfernandes/grpc-example/definitions"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Unable to listen: %s", err)
	}

	messages := make(chan int, 100)
	go func() {
		i := 0
		for {
			messages <- i
			time.Sleep(time.Second)
			i++
		}
	}()

	counter := NewCounter(messages)
	s := grpc.NewServer()
	definitions.RegisterCounterServer(s, counter)
	log.Printf("Starting server on %s", lis.Addr().String())
	s.Serve(lis)
}
