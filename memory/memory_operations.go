package memory

//These interfaces were created to interact with memory pages of 256 bytes

type reader interface {
	//Given and address it returns a byte
	Read(address byte) byte
}

type writer interface {
	//Given and address and a value it writes them on memory
	Write(address byte, value byte)
}

type reseter interface {
	//Sets the memory to 0 
	Reset()
}
