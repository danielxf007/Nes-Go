package main

import (
	"fmt"
	"nes_go/memory"
)

func main() {
	ram := new(memory.LinearMemory)
	ram.Data = make([]uint8, 5, 5)
	fmt.Println(ram.Read(0x00))
}