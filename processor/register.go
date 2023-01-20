package processor

type Register struct {
	Value byte
}

type ProgramCounter struct {
	ADH byte
	ADL byte
}

type FlagRegister struct {
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
	var borrow byte
	if pc.ADL < value {
	  borrow = 1
	}else {
	  borrow = 0
	}
	pc.ADL -= value
	pc.ADH -= borrow
	return uint16(borrow)
}

type FlagRegisterOperations interface {
	SetFlagC(bit byte)
	GetFlagC() byte
	SetFlagZ(bit byte)
	GetFlagZ() byte
	SetFlagI(bit byte)
	GetFlagI() byte
	SetFlagD(bit byte)
	GetFlagD() byte
	SetFlagB(bit byte)
	GetFlagB() byte
	SetFlagV(bit byte)
	GetFlagV() byte
	SetFlagN(bit byte)
	GetFlagN() byte
	Reset()
}

func (p* FlagRegister) SetFlagC(bit byte) {
	p.Value = SetNBit(p.Value, 0, bit)
}

func (p* FlagRegister) GetFlagC() byte {
	return GetNBit(p.Value, 0)
}

func (p* FlagRegister) SetFlagZ(bit byte) {
	p.Value = SetNBit(p.Value, 1, bit)
}

func (p* FlagRegister) GetFlagZ() byte {
	return GetNBit(p.Value, 1)
}

func (p* FlagRegister) SetFlagI(bit byte) {
	p.Value = SetNBit(p.Value, 2, bit)
}

func (p* FlagRegister) GetFlagI() byte {
	return GetNBit(p.Value, 2)
}

func (p* FlagRegister) SetFlagD(bit byte) {
	p.Value = SetNBit(p.Value, 3, bit)
}

func (p* FlagRegister) GetFlagD() byte {
	return GetNBit(p.Value, 3)
}

func (p* FlagRegister) SetFlagB(bit byte) {
	p.Value = SetNBit(p.Value, 4, bit)
}

func (p* FlagRegister) GetFlagB() byte {
	return GetNBit(p.Value, 4)
}

func (p* FlagRegister) SetFlagV(bit byte) {
	p.Value = SetNBit(p.Value, 6, bit)
}

func (p* FlagRegister) GetFlagV() byte {
	return GetNBit(p.Value, 6)
}

func (p* FlagRegister) SetFlagN(bit byte) {
	p.Value = SetNBit(p.Value, 7, bit)
}

func (p* FlagRegister) GetFlagN() byte {
	return GetNBit(p.Value, 7)
}

func (p* FlagRegister) Reset() {
	p.Value = 0x00
}
