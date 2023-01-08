package processor

func (addr* AddressBuffer) Increment(value byte) uint16 {
	carry := byte((uint16(addr.ADL) + uint16(value)) >> 8)
	addr.ADL += value
	addr.ADH += carry
	return uint16(carry)
}

//Boolean

func updateFlagsAND(cpu* CPU, result byte) {
	cpu.P.N = (result & 0x80) >> 7
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
}

func ExecuteAnd(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	cpu.A.Value &= value
	updateFlagsAND(cpu, cpu.A.Value)
	cpu.PC.Increment(1)
}

func ANDImmediate(cpu* CPU) uint16 {
  ExecuteAnd(cpu)
	return 2
}

func ANDZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteAnd(cpu)
	return 2
}

func ANDZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteAnd(cpu)
	return 2
}

func ANDAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteAnd(cpu)
	return 4
}

func ANDAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteAnd(cpu)
	return 4 + adjusted
}

func ANDAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteAnd(cpu)
	return 4 + adjusted
}

func ANDIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteAnd(cpu)
	return 6
}

func ANDIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteAnd(cpu)
	return 5 + adjusted
}

func (cpu* CPU) Execute(n_cycles uint16) {
	var current_cycles uint16 = 0
	for current_cycles < n_cycles {
		op_code := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
		cpu.PC.Increment(1)
		current_cycles += func_table[op_code](cpu)
	}
}
