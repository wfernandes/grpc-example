package main

import "github.com/wfernandes/grpc-example/definitions"

type Counter struct {
	messages chan int
}

func NewCounter(messages chan int) *Counter {
	return &Counter{
		messages: messages,
	}
}

func (s *Counter) Count(r *definitions.CountRequest, srv definitions.Counter_CountServer) error {
	// m := make(chan int)
	// go func() {
	// 	i := 0
	// 	for {
	// 		m <- i
	// 		time.Sleep(time.Second)
	// 		i++
	// 	}
	// }()

	for n := range s.messages {
		err := srv.Send(&definitions.CountResponse{
			Name:  r.Name,
			Count: int32(n),
		})
		if err != nil {
			// TODO: Handle error
		}

	}

	return nil
}
