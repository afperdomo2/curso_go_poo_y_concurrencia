package main

import (
	"fmt"

	afperdomo2Utils "github.com/afperdomo2/hellogomod"
	"github.com/donvito/hellomod"
	hellomodV2 "github.com/donvito/hellomod/v2"
)

func main() {
	fmt.Println("Hola, mundo!")

	hellomod.SayHello()

	hellomodV2.SayHello("🚀🚀🚀")

	afperdomo2Utils.HelloWorld()
}
