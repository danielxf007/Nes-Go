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
	SP Register
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
	//Arithmetic
	//ADC
	func_table[0x69] = ADCImmediate
	func_table[0x65] = ADCZeroPage
	func_table[0x75] = ADCZeroPageX
	func_table[0x6D] = ADCAbsolute
	func_table[0x7D] = ADCAbsoluteX
	func_table[0x79] = ADCAbsoluteY
	func_table[0x61] = ADCIndirectX
	func_table[0x71] = ADCIndirectY
	//INC
	func_table[0xE6] = INCZeroPage
	func_table[0xF6] = INCZeroPageX
	func_table[0xEE] = INCAbsolute
	func_table[0xFE] = INCAbsoluteX
	//INX
	func_table[0xE8] = INX
	//INY
	func_table[0xC8] = INY
	//DEC
	func_table[0xC6] = DECZeroPage
	func_table[0xD6] = DECZeroPageX
	func_table[0xCE] = DECAbsolute
	func_table[0xDE] = DECAbsoluteX
	//DEX
	func_table[0xCA] = DEX
	//DEY
	func_table[0x88] = DEY
	//Bitwise
	//ASL
	func_table[0x0A] = ASLA
	func_table[0x06] = ASLZeroPage
	func_table[0x16] = ASLZeroPageX
	func_table[0x0E] = ASLAbsolute
	func_table[0x1E] = ASLAbsoluteX
	//LSR
	func_table[0x4A] = LSRA
	func_table[0x46] = LSRZeroPage
	func_table[0x56] = LSRZeroPageX
	func_table[0x4E] = LSRAbsolute
	func_table[0x5E] = LSRAbsoluteX
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
	//EOR
	func_table[0x49] = EORImmediate
	func_table[0x45] = EORZeroPage
	func_table[0x55] = EORZeroPageX
	func_table[0x4D] = EORAbsolute
	func_table[0x5D] = EORAbsoluteX
	func_table[0x59] = EORAbsoluteY
	func_table[0x41] = EORIndirectX
	func_table[0x51] = EORIndirectY
	//Load
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
	//Store
	//STA
	func_table[0x85] = STAZeroPage
	func_table[0x95] = STAZeroPageX
	func_table[0x8D] = STAAbsolute
	func_table[0x9D] = STAAbsoluteX
	func_table[0x99] = STAAbsoluteY
	func_table[0x81] = STAIndirectX
	func_table[0x91] = STAIndirectY
	//STX
	func_table[0x86] = STXZeroPage
	func_table[0x96] = STXZeroPageY
	func_table[0x8E] = STXAbsolute
	//STY
	func_table[0x84] = STYZeroPage
	func_table[0x94] = STYZeroPageX
	func_table[0x8C] = STYAbsolute
	//Register Transfer
	//TAX
	func_table[0xAA] = TAX
	//TAY
	func_table[0xA8] = TAY
	//TXA
	func_table[0x8A] = TXA
	//Stack
	//TSX
	func_table[0xBA] = TSX
	//TXS
	func_table[0x9A] = TXS
	//PHA
	func_table[0x48] = PHA
	//PHP
	func_table[0x08] = PHP
	//PLA
	func_table[0x68] = PLA
	//PLP
	func_table[0x28] = PLP
	//Flag Operations
	//CLC
	func_table[0x18] = CLC
	//SEC
	func_table[0x38] = SEC
	//CLI
	func_table[0x58] = CLI
	//SEI
	func_table[0x78] = SEI
	//CLD
	func_table[0xD8] = CLD
	//SED
	func_table[0xF8] = SED
	//CLV
	func_table[0xB8] = CLV

}
