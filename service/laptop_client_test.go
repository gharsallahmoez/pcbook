package service_test

import (
	"context"
	"fmt"
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/gharsallahmoez/pcbook/sample"
	"github.com/gharsallahmoez/pcbook/service"
	"github.com/gharsallahmoez/pcbook/testdata"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	server, address := newTestingServer()
	// connect to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("error while creating connection %v", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	testLaptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: testLaptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		t.Errorf("expected success got error %v", err)
	}
	if res.Id != testLaptop.Id {
		t.Errorf("expected laptop with id %v got id %v", testLaptop.Id, res.Id)
	}

	// check if the laptop is successfully inserted
	insertedLaptop, err := server.Store.Find(res.Id)
	if err != nil {
		fmt.Errorf("error while getting laptop")
	}
	if insertedLaptop.Id != res.Id {
		t.Errorf("cannot get id")
	}
}

func TestLaptopServer_SearchLaptop(t *testing.T) {
	t.Parallel()
	server, address := newTestingServer()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server")
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()
		server.Store.Save(laptop)
	}
	for _, tc := range testdata.CreateSearchLaptopTT() {
		stream, err := laptopClient.SearchLaptop(context.Background(), tc.Request)
		if err != nil && !tc.HasError {
			t.Errorf("expected success got error %v", err)
		}
		if err == nil && tc.HasError {
			t.Error("expected error got success")
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil && !tc.HasError {
				t.Errorf("expected success got error %v", err)
			}
			if err == nil && tc.HasError {
				t.Error("expected error got success")
			}
			if res.Laptop.Id == "" && !tc.HasError {
				t.Errorf("expect getting at least one laptop id, got nil")
			}
		}
	}
}

func newTestingServer() (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewLaptopMemoryStore())
	server := grpc.NewServer()
	pb.RegisterLaptopServiceServer(server, laptopServer)
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Errorf("cannot create listener %v", err)
	}
	go server.Serve(lis)
	return laptopServer, lis.Addr().String()
}
