package testdata

import (
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/gharsallahmoez/pcbook/sample"
	"github.com/gharsallahmoez/pcbook/service"
	"google.golang.org/grpc/codes"
)

type AddLaptopTT struct {
	Name     string
	Request  *pb.CreateLaptopRequest
	Store    *service.LaptopMemoryStore
	Code     codes.Code
	HasError bool
}

type SearchLaptopTT struct {
	Name     string
	Request  *pb.SearchLaptopRequest
	Store    *service.LaptopMemoryStore
	Codes    codes.Code
	HasError bool
}

func CreateAddLaptopTT() []AddLaptopTT {
	// empty laptop id
	laptopWithoutID := sample.NewLaptop()
	laptopWithoutID.Id = ""
	// invalid laptop uuid
	laptopWithInvalidID := sample.NewLaptop()
	laptopWithInvalidID.Id = "invalid-uuid"
	// already exist laptop id
	store := service.NewLaptopMemoryStore()
	Existlaptop := sample.NewLaptop()
	store.Save(Existlaptop)

	tt := []AddLaptopTT{
		{
			Name:     "valid request",
			Request:  &pb.CreateLaptopRequest{Laptop: sample.NewLaptop()},
			Store:    service.NewLaptopMemoryStore(),
			Code:     codes.OK,
			HasError: false,
		}, {
			Name:     "valid request without id",
			Request:  &pb.CreateLaptopRequest{Laptop: laptopWithoutID},
			Store:    service.NewLaptopMemoryStore(),
			Code:     codes.OK,
			HasError: false,
		}, {
			Name:     "invalid uuid",
			Request:  &pb.CreateLaptopRequest{Laptop: laptopWithInvalidID},
			Store:    service.NewLaptopMemoryStore(),
			Code:     codes.InvalidArgument,
			HasError: true,
		}, {
			Name:     "laptop id already exist",
			Request:  &pb.CreateLaptopRequest{Laptop: Existlaptop},
			Store:    store,
			Code:     codes.AlreadyExists,
			HasError: true,
		},
	}
	return tt
}

func CreateSearchLaptopTT() []SearchLaptopTT {
	tt := []SearchLaptopTT{
		{
			Name: "valid search request",
			Request: &pb.SearchLaptopRequest{
				Filter: &pb.Filter{
					MaxPriceUsd: 2000,
					MinCpuCores: 4,
					MinCpuGhz:   2.5,
					MinRam: &pb.Memory{
						Value: 2,
						Unit:  pb.Memory_GIGABYTE,
					},
				},
			},
			Store:    service.NewLaptopMemoryStore(),
			Codes:    codes.OK,
			HasError: false,
		},
	}
	return tt
}
