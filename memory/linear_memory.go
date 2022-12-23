package memory

//This type of memory only holds bytes and is volatile
type LinearMemory struct {
	Data []byte
}

func (memory* LinearMemory) Read(address byte) byte {
	return memory.Data[address]
}

func (memory* LinearMemory) Write(address byte, value byte) {
	memory.Data[address] = value
}

func (memory* LinearMemory) Reset() {
	for i := range memory.Data {
		memory.Data[i] = 0
	}
}

