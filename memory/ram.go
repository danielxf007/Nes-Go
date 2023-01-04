package memory

//This type of memory only holds bytes and is volatile
type Ram struct {
	Data [0x2000]byte
}

func (ram* Ram) Read(address uint16) byte {
	return ram.Data[address % uint16(len(ram.Data))]
}

func (ram* Ram) Write(address uint16, value byte) {
	ram.Data[address] = value
}

func (ram* Ram) Reset() {
	for i := range ram.Data {
		ram.Data[i] = 0
	}
}

