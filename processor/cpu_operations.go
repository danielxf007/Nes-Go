package processor

func (addr* AddressBuffer) Increment(value byte) uint16 {
	carry := byte((uint16(addr.ADL) + uint16(value)) >> 8)
	addr.ADL += value
	addr.ADH += carry
	return uint16(carry)
}

//Arithmetic

//Shift Left
func updateFlagsASL(cpu* CPU, bit byte, result byte) {
	cpu.P.C = bit
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
}

func ExecuteASLA(cpu* CPU) {
	bit := GetNBit(cpu.A.Value, 7)
	cpu.A.Value <<= 1
	updateFlagsASL(cpu, bit, cpu.A.Value)
}

func ExecuteASLMEM(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	bit := GetNBit(value, 7)
	value <<= 1
	updateFlagsASL(cpu, bit, value)
	cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, value)
	cpu.PC.Increment(1)
}

func ASLA(cpu* CPU) uint16 {
	ExecuteASLA(cpu)
	return 1
}

func ASLZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteASLMEM(cpu)
	return 5
}

func ASLZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteASLMEM(cpu)
	return 6
}

func ASLAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteASLMEM(cpu)
	return 6
}

func ASLAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteASLMEM(cpu)
	return 7
}

//Boolean

//AND
func updateFlagsAND(cpu* CPU, result byte) {
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
}

func ExecuteAND(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	cpu.A.Value &= value
	updateFlagsAND(cpu, cpu.A.Value)
	cpu.PC.Increment(1)
}

func ANDImmediate(cpu* CPU) uint16 {
  ExecuteAND(cpu)
	return 2
}

func ANDZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteAND(cpu)
	return 2
}

func ANDZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteAND(cpu)
	return 2
}

func ANDAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteAND(cpu)
	return 4
}

func ANDAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteAND(cpu)
	return 4 + adjusted
}

func ANDAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteAND(cpu)
	return 4 + adjusted
}

func ANDIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteAND(cpu)
	return 6
}

func ANDIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteAND(cpu)
	return 5 + adjusted
}

//ORA
func updateFlagsORA(cpu* CPU, result byte) {
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
}

func ExecuteORA(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	cpu.A.Value |= value
	updateFlagsORA(cpu, cpu.A.Value)
	cpu.PC.Increment(1)
}

func ORAImmediate(cpu* CPU) uint16 {
  ExecuteORA(cpu)
	return 2
}

func ORAZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteORA(cpu)
	return 2
}

func ORAZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteORA(cpu)
	return 2
}

func ORAAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteORA(cpu)
	return 4
}

func ORAAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteORA(cpu)
	return 4 + adjusted
}

func ORAAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteORA(cpu)
	return 4 + adjusted
}

func ORAIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteORA(cpu)
	return 6
}

func ORAIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteORA(cpu)
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
