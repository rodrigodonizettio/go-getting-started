package sample

import (
	"math/rand"
	"time"

	pb "br.com.rodrigodonizettio/grpc-pcbook/pb/proto"
	"github.com/google/uuid"
)

//This function will always be called before the others and will change the seed to always generate different values
func init() {
	rand.Seed(time.Now().UnixNano())
}

//   Name(params)           returnType {
func RandomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	case 3:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func RandomLaptopName(brand string) string {
	switch brand {
	case "Asus":
		return RandomStringFromSet("ROG", "VivoBook", "ZenBook")
	case "Dell":
		return RandomStringFromSet("Alienware", "Inspiron", "Vostro")
	default:
		return RandomStringFromSet("IdeaPad", "Legion", "ThinkPad")
	}
}

func RandomLaptopBrand() string {
	return RandomStringFromSet("Asus", "Dell", "Lenovo")
}

func RandomUuid() string {
	return uuid.New().String()
}

func RandomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func RandomScreenResolution() *pb.Screen_Resolution {
	height := RandomInt(800, 5000)
	width := height * (16 / 9)

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

func RandomGpuBrand() string {
	return RandomStringFromSet("NVidia", "AMD")
}

func RandomGpuName(brand string) string {
	if brand == "NVidia" {
		return RandomStringFromSet("930", "1080", "2060")
	}
	return RandomStringFromSet("580, 590, Vega")
}

func RandomCpuBrand() string {
	return RandomStringFromSet("Intel", "AMD")
}

func RandomCpuName(brand string) string {
	if brand == "Intel" {
		return RandomStringFromSet("Xeon", "Core i3", "Core i5", "Core i7", "Core i9")
	}
	return RandomStringFromSet("Ryzen 3", "Ryzen 5", "Ryzen 7")
}

//Receives a stringList and return a random single element of that
func RandomStringFromSet(stringList ...string) string {
	n := len(stringList)
	if n == 0 {
		return ""
	}
	return stringList[rand.Intn(n)]
}

func RandomBool() bool {
	//rand.Intn(n) generates values from 0 to < n
	return rand.Intn(2) == 1
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
