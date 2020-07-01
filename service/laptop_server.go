package service

import (
	"context"
	"errors"
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		store,
	}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	// check the context
	//time.Sleep(3 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		log.Print("context deadline exceeded")
		return nil, status.Errorf(codes.DeadlineExceeded, "context deadline exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Print("context is canceled")
		return nil, status.Errorf(codes.Canceled, "context canceled")
	}
	laptop := req.GetLaptop()
	if len(laptop.Id) > 0 {
		// check if it's a real uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot parse laptop id %v", err)
		}
	} else {
		id := uuid.New()
		laptop.Id = id.String()
	}
	// save the laptop
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, LaptopExist) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store %v", err)
	}
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	log.Printf("laptop saved successfully with this id %v", laptop.Id)
	return res, nil
}

func (server *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	// the call back function
	f := func(laptop *pb.Laptop) error {
		res := &pb.SearchLaptopResponse{
			Laptop: laptop,
		}
		err := stream.Send(res)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot send stream %v", err)
		}
		log.Printf("send laptop with id %v", laptop.Id)

		return nil
	}
	// call search function by passing the callback function f
	err := server.Store.Search(stream.Context(), filter, f)
	if err != nil {
		log.Printf("error %v", err)
		return status.Errorf(codes.Internal, "no laptop available %v", err)
	}
	return nil
}
