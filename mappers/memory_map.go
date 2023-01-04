package mappers
import (
	"nes_go/memory"
)

type MemoryMap struct {
	Ram *memory.Ram
	Rom *memory.Rom
	Memory_pages [256]memory.MemoryOperations
}

var Memory_map MemoryMap 

func initMemoryPages() {
	for index := 0; index < 8; index++ {
		Memory_map.Memory_pages[index] = Memory_map.Ram
		Memory_map.Memory_pages[index+8] = Memory_map.Ram
		Memory_map.Memory_pages[index+16] = Memory_map.Ram
	}
	for index := 223; index < 256; index++ {
		Memory_map.Memory_pages[index] = Memory_map.Rom
	}
}

func init () {
	Memory_map = MemoryMap{Ram: new(memory.Ram), Rom: new(memory.Rom),}
	initMemoryPages()
}