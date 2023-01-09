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
	//ASL
	func_table[0x0A] = ASLA
	func_table[0x06] = ASLZeroPage
	func_table[0x16] = ASLZeroPageX
	func_table[0x0E] = ASLAbsolute
	func_table[0x1E] = ASLAbsoluteX
	//AND
	func_table[0x29] = ANDImmediate
	func_table[0x25] = ANDZeroPage
	func_table[0x35] = ANDZeroPageX
	func_table[0x2D] = ANDAbsolute
	func_table[0x3D] = ANDAbsoluteX
	func_table[0x39] = ANDAbsoluteY
	func_table[0x21] = ANDIndirectX
	func_table[0x31] = ANDIndirectY
	//ORA
	func_table[0x09] = ORAImmediate
	func_table[0x05] = ORAZeroPage
	func_table[0x15] = ORAZeroPageX
	func_table[0x0D] = ORAAbsolute
	func_table[0x1D] = ORAAbsoluteX
	func_table[0x19] = ORAAbsoluteY
	func_table[0x01] = ORAIndirectX
	func_table[0x11] = ORAIndirectY
	//LDA
	func_table[0xA9] = LDAImmediate
	func_table[0xA5] = LDAZeroPage
	func_table[0xB5] = LDAZeroPageX
	func_table[0xAD] = LDAAbsolute
	func_table[0xBD] = LDAAbsoluteX
	func_table[0xB9] = LDAAbsoluteY
	func_table[0xA1] = LDAIndirectX
	func_table[0xB1] = LDAIndirectY
	//LDX
	func_table[0xA2] = LDXImmediate
	func_table[0xA6] = LDXZeroPage
	func_table[0xB6] = LDXZeroPageY
	func_table[0xAE] = LDXAbsolute
	func_table[0xBE] = LDXAbsoluteY
	//LDY
	func_table[0xA0] = LDYImmediate
	func_table[0xA4] = LDYZeroPage
	func_table[0xB4] = LDYZeroPageX
	func_table[0xAC] = LDYAbsolute
	func_table[0xBC] = LDYAbsoluteX
}
