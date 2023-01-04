package mappers

type NoMapper struct {
}

func (mapper* NoMapper) Read(page_number byte, offset byte) byte {
	address := (uint16(page_number) << 8) | uint16(offset)
	return Memory_map.Memory_pages[page_number].Read(address)
}

func (mapper* NoMapper) Write(page_number byte, offset byte, value byte) {
	address := (uint16(page_number) << 8) | uint16(offset)
	Memory_map.Memory_pages[page_number].Write(address, value)
}