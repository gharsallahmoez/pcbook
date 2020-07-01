package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gharsallahmoez/pcbook/pb"
	"github.com/jinzhu/copier"
	"sync"
	"time"
)

var LaptopExist = errors.New("laptop already exist")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

type LaptopMemoryStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewLaptopMemoryStore() *LaptopMemoryStore {
	return &LaptopMemoryStore{
		data: make(map[string]*pb.Laptop),
	}
}

// 53 93 47 79
// 450
// 30
func (store *LaptopMemoryStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[laptop.Id] != nil {
		return LaptopExist
	}
	l := &pb.Laptop{}
	err := copier.Copy(l, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop object %v", err)
	}
	store.data[laptop.Id] = l
	return nil
}

func (store *LaptopMemoryStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptop := store.data[id]
	if laptop == nil {
		return nil, fmt.Errorf("laptop non exist")
	}
	cLaptop := &pb.Laptop{}
	err := copier.Copy(cLaptop, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop %v", err)
	}
	return laptop, nil
}

func (store *LaptopMemoryStore) Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	for _, laptop := range store.data {
		time.Sleep(time.Second)
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			return ctx.Err()
		}
		if isQualifier(filter, laptop) {
			other := &pb.Laptop{}
			err := copier.Copy(other, laptop)
			if err != nil {
				return err
			}
			err = found(laptop)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualifier(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.PriceUsd > filter.MaxPriceUsd {
		return false
	}
	if laptop.Cpu.NumberCores < filter.MinCpuCores {
		return false
	}
	if laptop.Cpu.MinGhz < filter.MinCpuGhz {
		return false
	}
	if toBit(laptop.Ram) < toBit(filter.MinRam) {
		return false
	}
	return true
}

func toBit(ram *pb.Memory) uint64 {
	value := ram.GetValue()

	switch ram.Unit {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3 // 1 byte = 8 bit and 8 = 2^3
	case pb.Memory_KILOBYTE:
		return value << 13 // 1 kb = 1024* 8 bit = 2¹0 * 2^3 = 2¹3
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}
