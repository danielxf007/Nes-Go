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
	Value byte
}

type ProgramCounterOperations interface {
	SetADH(value byte)
	SetADL(value byte)
	//Increments the PC by value and returns 1 or 0 if ADL overflowed
	Increment(value byte) uint16
	//Increments the PC by value and returns 1 or 0 if ADL underflowed
	Decrement(value byte) uint16
}

func (pc* ProgramCounter) Increment(value byte) uint16 {
	carry := byte((uint16(pc.ADL) + uint16(value)) >> 8)
	pc.ADL += value
	pc.ADH += carry
	return uint16(carry)
}

func (pc* ProgramCounter) Decrement(value byte) uint16 {
	return 0
}

type FlagRegisterOperations interface {
  UpdateValue()
  SetValue(value byte)
	Reset()
}

func (p* FlagRegister) UpdateValue() {
  p.Value = (p.N << 7)|(p.V << 6)|(p.B << 4)|(p.D << 3)|(p.I << 2)|(p.Z << 1)|(p.C)
}

func (p* FlagRegister) SetValue(value byte) {
  p.Value = value
  p.C, p.Z, p.I, p.D = GetNBit(value, 0), GetNBit(value, 1), GetNBit(value, 2), GetNBit(value, 3)
  p.B, p.V, p.N = GetNBit(value, 4), GetNBit(value, 6), GetNBit(value, 7)
}

func (p* FlagRegister) Reset() {
	p.C, p.Z, p.I, p.D, p.B, p.V, p.N = 0, 0, 0, 0, 0, 0, 0
	p.UpdateValue()
}
