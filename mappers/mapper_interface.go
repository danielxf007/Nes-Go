package mappers

type MapperInterface interface {
	Read(page_number byte, offset byte) byte
	Write(page_number byte, offset byte, value byte)
}