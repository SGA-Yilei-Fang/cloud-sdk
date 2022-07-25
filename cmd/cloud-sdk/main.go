package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SGA-Bicheng-Zhang/cloud-sdk/server"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	fmt.Println("Start listening localhost:9001")

	server := server.NewServer()
	server.Serve(listen)
}
