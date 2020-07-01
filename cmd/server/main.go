package main

import (
	"flag"
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/gharsallahmoez/pcbook/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// get address from command line
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	// instantiate the laptop server
	laptopServer := service.NewLaptopServer(service.NewLaptopMemoryStore())
	// create grpc server
	grpcServer := grpc.NewServer()
	// register the server
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	// create listener
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("cannot create listener %v", err)
	}
	// serve
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("cannot serve %v", err)
	}
}
