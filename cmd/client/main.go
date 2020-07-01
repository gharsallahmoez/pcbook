package main

import (
	"context"
	"flag"
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/gharsallahmoez/pcbook/sample"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address that we will connect to")
	flag.Parse()
	log.Printf("connecting to server %v", *serverAddress)
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to the client %v", err)
		return
	}
	client := pb.NewLaptopServiceClient(conn)
	addLaptop(client)
	searchLaptop(client)
}

func addLaptop(client pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		log.Fatalf("cannot add laptop %v", err)
		return
	}
	log.Printf("laptop created with id %v", res.Id)
}

func searchLaptop(client pb.LaptopServiceClient) {
	// create 10 laptops
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 10; i++ {
		res, err := client.CreateLaptop(ctx, &pb.CreateLaptopRequest{
			Laptop: sample.NewLaptop(),
		})
		if err != nil {
			log.Fatalf("cannot add laptop %v", err)
			return
		}
		log.Printf("laptop created with id %v", res.Id)
	}
	// search for laptop
	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 8,
		MinCpuGhz:   2.6,
		MinRam: &pb.Memory{
			Value: 2,
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	SearchReq := &pb.SearchLaptopRequest{
		Filter: filter,
	}
	stream, err := client.SearchLaptop(ctx, SearchReq)
	if err != nil {
		log.Fatalf("cannot search laptops %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot read stream %v", err)
		}
		//time.Sleep(time.Second)
		log.Printf("found laptop with id : %v \n", res.GetLaptop().GetId())
	}
}
