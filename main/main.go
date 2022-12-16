package main

import (
	"fmt"
	"memory"
)

func main() {

	ram := new(linearMemory)
	ram.data = make([]uint8, 0x2000, 0x2000)
	fmt.Println(ram)
}