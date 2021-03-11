package sample

import (
	pb "br.com.rodrigodonizettio/grpc-pcbook/pb/proto"
	"github.com/golang/protobuf/ptypes"
)

//   Name(params)  returnType {
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  RandomKeyboardLayout(),
		Backlit: RandomBool(),
	}
	return keyboard
}

func NewCpu() *pb.Cpu {
	brand := RandomCpuBrand()
	name := RandomCpuName(brand)
	numberCores := RandomInt(2, 8)
	numberThreads := RandomInt(numberCores, 12)
	minGhz := RandomFloat64(1.0, 4.5)
	maxGhz := RandomFloat64(minGhz, 5.5)

	cpu := &pb.Cpu{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGpu() *pb.Gpu {
	brand := RandomGpuBrand()
	name := RandomGpuName(brand)
	minGhz := RandomFloat64(2.0, 5.5)
	maxGhz := RandomFloat64(minGhz, 6.5)
	memory := &pb.Memory{
		Value: uint64(RandomInt(1, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.Gpu{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(RandomInt(8, 32)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

func NewSsd() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(RandomInt(128, 2048)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func NewHdd() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(RandomInt(1, 5)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   RandomFloat32(14, 28),
		Resolution: RandomScreenResolution(),
		Panel:      RandomScreenPanel(),
		Multitouch: RandomBool(),
	}
	return screen
}

func NewLaptop() *pb.Laptop {
	brand := RandomLaptopBrand()
	name := RandomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       RandomUuid(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCpu(),
		Ram:      NewRam(),
		Gpus:     []*pb.Gpu{NewGpu()},
		Storages: []*pb.Storage{NewSsd(), NewHdd()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: RandomFloat64(0.5, 2.5),
		},
		PriceBrl:    RandomFloat64(4000, 12000),
		ReleaseYear: uint32(RandomInt(2015, 2021)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
