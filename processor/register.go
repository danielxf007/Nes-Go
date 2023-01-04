package processor

type Register struct {
	Value byte
}

type ProgramCounter struct {
	ADH byte
	ADL byte
}

type FlagRegister struct {
	C byte
	Z byte
	I byte
	D byte
	B byte
	V byte
	N byte
}

type ProgramCounterOperations interface {
	SetADH(value byte)
	SetADL(value byte)
	//Increments the PC by value and returns 1 or 0 if ADL overflowed
	Increment(value byte) byte
	//Increments the PC by value and returns 1 or 0 if ADL underflowed
	Decrement(value byte) byte
}

func (pc* ProgramCounter) Increment(value byte) byte {
	carry := byte((uint16(pc.ADL) + uint16(value)) >> 8)
	pc.ADL += value
	pc.ADH += carry
	return carry
}

func (pc* ProgramCounter) Decrement(value byte) byte {
	return 0
}

type FlagRegisterOperations interface {
	Reset()
}

func (p* FlagRegister) Reset() {
	p.C, p.Z, p.I, p.D, p.B, p.V, p.N = 0, 0, 0, 0, 0, 0, 0
}