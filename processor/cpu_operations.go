package processor

func (addr* AddressBuffer) Increment(value byte) uint16 {
	carry := byte((uint16(addr.ADL) + uint16(value)) >> 8)
	addr.ADL += value
	addr.ADH += carry
	return uint16(carry)
}

//Arithmetic

//INC
func updateFlagsINC(cpu* CPU, result byte) {
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
}

func ExecuteINCReg(cpu* CPU, reg* Register) {
	reg.Value++
	updateFlagsINC(cpu, reg.Value)
}

func ExecuteINCMEM(cpu* CPU) {
  value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	value++
	cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, value)
	updateFlagsINC(cpu, value)
	cpu.PC.Increment(1)
}

//INC
func INCZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteINCMEM(cpu)
	return 5
}

func INCZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteINCMEM(cpu)
	return 6
}

func INCAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteINCMEM(cpu)
	return 6
}

func INCAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteINCMEM(cpu)
	return 7
}

//INX
func INX(cpu* CPU) uint16 {
  ExecuteINCReg(cpu, &cpu.X)
	return 2
}

//INY
func INY(cpu* CPU) uint16 {
  ExecuteINCReg(cpu, &cpu.Y)
	return 2
}

//Bitwise

//ASL
func updateFlagsASL(cpu* CPU, bit byte, result byte) {
	cpu.P.C = bit
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
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
	return 2
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

//LSR
func updateFlagsLSR(cpu* CPU, bit byte, result byte) {
	cpu.P.C = bit
	cpu.P.N = 0
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
}

func ExecuteLSRA(cpu* CPU) {
	bit := GetNBit(cpu.A.Value, 0)
	cpu.A.Value >>= 1
	updateFlagsLSR(cpu, bit, cpu.A.Value)
}

func ExecuteLSRMEM(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	bit := GetNBit(value, 0)
	value >>= 1
	updateFlagsASL(cpu, bit, value)
	cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, value)
	cpu.PC.Increment(1)
}

func LSRA(cpu* CPU) uint16 {
	ExecuteLSRA(cpu)
	return 2
}

func LSRZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteLSRMEM(cpu)
	return 5
}

func LSRZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteLSRMEM(cpu)
	return 6
}

func LSRAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteLSRMEM(cpu)
	return 6
}

func LSRAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteLSRMEM(cpu)
	return 7
}

//AND
func updateFlagsAND(cpu* CPU, result byte) {
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
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
	return 3
}

func ANDZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteAND(cpu)
	return 4
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
	cpu.P.UpdateValue()
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
	return 3
}

func ORAZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteORA(cpu)
	return 4
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

//EOR
func updateFlagsEOR(cpu* CPU, result byte) {
	cpu.P.N = GetNBit(result, 7)
	if result == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
}

func ExecuteEOR(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	cpu.A.Value ^= value
	updateFlagsORA(cpu, cpu.A.Value)
	cpu.PC.Increment(1)
}

func EORImmediate(cpu* CPU) uint16 {
  ExecuteEOR(cpu)
	return 2
}

func EORZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteEOR(cpu)
	return 3
}

func EORZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteEOR(cpu)
	return 4
}

func EORAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteEOR(cpu)
	return 4
}

func EORAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteEOR(cpu)
	return 4 + adjusted
}

func EORAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteEOR(cpu)
	return 4 + adjusted
}

func EORIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteEOR(cpu)
	return 6
}

func EORIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteEOR(cpu)
	return 5 + adjusted
}

//Load Registers

func updateFlagsLoadRegister(cpu* CPU, value byte) {
	cpu.P.N = GetNBit(value, 7)
	if value == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
}

func ExecuteLoadRegister(cpu* CPU, register* Register) {
  register.Value = cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
  updateFlagsLoadRegister(cpu, register.Value)
  cpu.PC.Increment(1)
}

//LDA
func LDAImmediate(cpu* CPU) uint16 {
  ExecuteLoadRegister(cpu, &cpu.A)
	return 2
}

func LDAZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 3
}

func LDAZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 4
}

func LDAAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 4
}

func LDAAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 4 + adjusted
}

func LDAAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 4 + adjusted
}

func LDAIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 6
}

func LDAIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.A)
	return 5 + adjusted
}

//LDX
func LDXImmediate(cpu* CPU) uint16 {
  ExecuteLoadRegister(cpu, &cpu.X)
	return 2
}

func LDXZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.X)
	return 3
}

func LDXZeroPageY(cpu* CPU) uint16 {
	GetZeroPageYAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.X)
	return 4
}

func LDXAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.X)
	return 4
}

func LDXAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.X)
	return 4 + adjusted
}

//LDY
func LDYImmediate(cpu* CPU) uint16 {
  ExecuteLoadRegister(cpu, &cpu.Y)
	return 2
}

func LDYZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.Y)
	return 3
}

func LDYZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.Y)
	return 4
}

func LDYAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.Y)
	return 4
}

func LDYAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteLoadRegister(cpu, &cpu.Y)
	return 4 + adjusted
}

//Store Registers
func ExecuteStoreRegister(cpu* CPU, register* Register) {
  cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, register.Value)
  cpu.PC.Increment(1)
}

//STA
func STAZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 3
}

func STAZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 4
}

func STAAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 4
}

func STAAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 5
}

func STAAbsoluteY(cpu* CPU) uint16 {
	GetAbsoluteYAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 5
}

func STAIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 6
}

func STAIndirectY(cpu* CPU) uint16 {
	GetIndirectYAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.A)
	return 6
}

//STX
func STXZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.X)
	return 3
}

func STXZeroPageY(cpu* CPU) uint16 {
	GetZeroPageYAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.X)
	return 4
}

func STXAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.X)
	return 4
}

//STY
func STYZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.Y)
	return 3
}

func STYZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.Y)
	return 4
}

func STYAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteStoreRegister(cpu, &cpu.Y)
	return 4
}

//Transfer value among registers
func updateFlagsTransferValueRegisters(cpu* CPU, value byte) {
	cpu.P.N = GetNBit(value, 7)
	if value == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
}


func ExecuteTransferValueRegisters(cpu* CPU, src_reg* Register, target_reg* Register) {
  target_reg.Value = src_reg.Value
  updateFlagsTransferValueRegisters(cpu, src_reg.Value)
}

//TAX
func TAX(cpu* CPU) uint16 {
  ExecuteTransferValueRegisters(cpu, &cpu.A, &cpu.X)
  return 2
}

//TAY
func TAY(cpu* CPU) uint16 {
  ExecuteTransferValueRegisters(cpu, &cpu.A, &cpu.Y)
  return 2
}

//TXA
func TXA(cpu* CPU) uint16 {
  ExecuteTransferValueRegisters(cpu, &cpu.X, &cpu.A)
  return 2
}

//TYA
func TYA(cpu* CPU) uint16 {
  ExecuteTransferValueRegisters(cpu, &cpu.Y, &cpu.A)
  return 2
}

//Stack instructions

//TXS
func TXS(cpu* CPU) uint16 {
  cpu.SP.Value = cpu.X.Value
  return 2
}

//TSX
func TSX(cpu* CPU) uint16 {
  ExecuteTransferValueRegisters(cpu, &cpu.SP, &cpu.X)
  return 2
}

//Push operations
func ExecutePush(cpu* CPU, value byte) {
  cpu.Mapper.Write(0x01, cpu.SP.Value, value)
  cpu.SP.Value--
}

//PHA
func PHA(cpu* CPU) uint16 {
  ExecutePush(cpu, cpu.A.Value)
  return 3
}

//PHP
func PHP(cpu* CPU) uint16 {
  ExecutePush(cpu, cpu.P.Value)
  return 3
}

//Pull operations
func updateFlagsPull(cpu* CPU, value byte) {
	cpu.P.N = GetNBit(value, 7)
	if value == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
	cpu.P.UpdateValue()
}

func ExecutePullA(cpu* CPU) {
  cpu.SP.Value++
  cpu.A.Value = cpu.Mapper.Read(0x01, cpu.SP.Value)
  updateFlagsPull(cpu, cpu.A.Value)
}

func ExecutePullP(cpu* CPU) {
  cpu.SP.Value++
  cpu.P.SetValue(cpu.Mapper.Read(0x01, cpu.SP.Value))
}

//PLA
func PLA(cpu* CPU) uint16 {
  ExecutePullA(cpu)
  return 4
}

//PLP
func PLP(cpu* CPU) uint16 {
  ExecutePullP(cpu)
  return 4
}


func (cpu* CPU) Execute(n_cycles uint16) {
	var current_cycles uint16 = 0
	for current_cycles < n_cycles {
		op_code := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
		cpu.PC.Increment(1)
		current_cycles += func_table[op_code](cpu)
	}
}
