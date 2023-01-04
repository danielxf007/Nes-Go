package memory

//These interfaces were created to interact with memory pages of 256 bytes

type MemoryOperations interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
	Reset()
}


