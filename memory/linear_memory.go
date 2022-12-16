package memory

type linearMemory struct {
	data uint8[]
}

type reader interface {
	read(address uint8) (uint8)
}

func (memory *linearMemory) read(address uint8) (uint8) {
	return 0x00
}