package service_test

import (
	"context"
	"github.com/gharsallahmoez/pcbook/service"
	"github.com/gharsallahmoez/pcbook/testdata"
	"google.golang.org/grpc/status"
	"testing"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()
	tt := testdata.CreateAddLaptopTT()
	for _, tc := range tt {
		store := service.NewLaptopServer(tc.Store)
		res, err := store.CreateLaptop(context.Background(), tc.Request)
		if tc.HasError && err == nil {
			t.Errorf("expected error got nil")
		}
		if !tc.HasError && err != nil {
			t.Errorf("expected success got error %v", err)
		}
		status, ok := status.FromError(err)
		if !ok {
			t.Errorf("cannot get error code")
		}
		if tc.HasError && err != nil && status.Code() != tc.Code {
			t.Errorf("expected error code %v got %v", tc.Code, status.Code())
		}
		if !tc.HasError && err == nil && res.Id == "" {
			t.Errorf("expected valid response got empty")
		}
	}
}
