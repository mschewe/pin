package main

import (
	"flag"
	"fmt"

	"github.com/mschewe/pin/lib"
)

var length, count int

func init() {
	flag.IntVar(&length, "n", 4, "length of the pin")
	flag.IntVar(&count, "c", 1, "number of generated pins")
	flag.Parse()
}

func main() {
	for i := 0; i < count; i++ {
		fmt.Println(pin.Generate(length))
	}
}
