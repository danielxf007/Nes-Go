package processor
import (
	"nes_go/mappers"
)

type AddressBuffer struct {
	ADH byte
	ADL byte
}
 
type CPU struct {
	A Register
	X Register
	Y Register
	P FlagRegister
	PC ProgramCounter
	Addr AddressBuffer
	Mapper mappers.MapperInterface	
	Halted bool
}

type Emulation interface {
	Execute(n_cycles uint32)
}

type AddressBufferOperations interface {
	Increment(value byte) uint16
}

var func_table [0xFF]func(cpu* CPU) uint16

func init() {
	func_table = [0xFF]func(cpu* CPU) uint16{}
	func_table[0x29] = ANDImmediate
}
