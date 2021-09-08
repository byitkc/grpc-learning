package main

import (
	"fmt"
	"github.com/byitkc/grpc-learning/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, I'm a client!")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := greetpb.NewGreetServiceClient(conn)
	log.Printf("Created client: %f", c)
}
