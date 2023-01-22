package processor

//NOP
func NOP(cpu* CPU) uint16 {
  return 2
}
//Arithmetic

//ADC
func updateFlagsADC(cpu* CPU, carry_bit byte, A_val_i byte, A_val_f byte) {
	cpu.P.SetFlagC(carry_bit)
	if A_val_f == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	if GetNBit(A_val_i, 7) != GetNBit(A_val_f, 7) {
	  cpu.P.SetFlagV(1)
	}else {
	  cpu.P.SetFlagV(0)
	}
	cpu.P.SetFlagN(GetNBit(A_val_f, 7))
}

func ExecuteADC(cpu* CPU) {
  value := uint16(cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL))
  value += (uint16(cpu.A.Value) + uint16(cpu.P.GetFlagC()))
  carry_bit := byte(value >> 8)
  A_val_i := cpu.A.Value
  A_val_f := byte(value)
  cpu.A.Value = A_val_f
  updateFlagsADC(cpu, carry_bit, A_val_i, A_val_f)
  cpu.PC.Increment(1) 
}

func ADCImmediate(cpu* CPU) uint16 {
  ExecuteADC(cpu)
	return 2
}

func ADCZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteADC(cpu)
	return 3
}

func ADCZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteADC(cpu)
	return 4
}

func ADCAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteADC(cpu)
	return 4
}

func ADCAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteADC(cpu)
	return 4 + adjusted
}

func ADCAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteADC(cpu)
	return 4 + adjusted
}

func ADCIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteADC(cpu)
	return 6
}

func ADCIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteADC(cpu)
	return 5 + adjusted
}

//SBC
func updateFlagsSBC(cpu* CPU) {
  if int8(cpu.A.Value) >= 0x00 { //There could be issues with this flag.
    cpu.P.SetFlagC(1)
  }else {
    cpu.P.SetFlagC(0)
  }
	if cpu.A.Value == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	if cpu.A.Value > 127 ||  int8(cpu.A.Value) < -128 {
		cpu.P.SetFlagV(1)
	}else {
		cpu.P.SetFlagV(0)
	}
	cpu.P.SetFlagN(GetNBit(cpu.A.Value, 7))
}

func ExecuteSBC(cpu* CPU) {
  var complemented_c_flag byte
  if cpu.P.GetFlagC() == 0x00 {
    complemented_c_flag = 0x01
  }else {
    complemented_c_flag = 0x00
  } 
  value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
  cpu.A.Value = cpu.A.Value - value - complemented_c_flag
  updateFlagsSBC(cpu)
  cpu.PC.Increment(1) 
}

func SBCImmediate(cpu* CPU) uint16 {
  ExecuteSBC(cpu)
	return 2
}

func SBCZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteSBC(cpu)
	return 3
}

func SBCZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteSBC(cpu)
	return 4
}

func SBCAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteSBC(cpu)
	return 4
}

func SBCAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteSBC(cpu)
	return 4 + adjusted
}

func SBCAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteSBC(cpu)
	return 4 + adjusted
}

func SBCIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteSBC(cpu)
	return 6
}

func SBCIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteSBC(cpu)
	return 5 + adjusted
}

//INCREMENT
func updateFlagsINC(cpu* CPU, result byte) {
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
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

//DECREMENT
func updateFlagsDEC(cpu* CPU, result byte) {
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
}

func ExecuteDECReg(cpu* CPU, reg* Register) {
	reg.Value--
	updateFlagsINC(cpu, reg.Value)
}

func ExecuteDECMEM(cpu* CPU) {
  value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	value--
	cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, value)
	updateFlagsINC(cpu, value)
	cpu.PC.Increment(1)
}

//DEC
func DECZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteDECMEM(cpu)
	return 5
}

func DECZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteDECMEM(cpu)
	return 6
}

func DECAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteDECMEM(cpu)
	return 6
}

func DECAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteDECMEM(cpu)
	return 7
}

//INX
func DEX(cpu* CPU) uint16 {
  ExecuteDECReg(cpu, &cpu.X)
	return 2
}

//INY
func DEY(cpu* CPU) uint16 {
  ExecuteDECReg(cpu, &cpu.Y)
	return 2
}

//Bitwise

//Shifts

//ASL
func updateFlagsASL(cpu* CPU, bit byte, result byte) {
	cpu.P.SetFlagC(bit)
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
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
	cpu.P.SetFlagC(bit)
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
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

//ROL
func updateFlagsROL(cpu* CPU, bit byte, result byte) {
	cpu.P.SetFlagC(bit)
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
}

func ExecuteROLA(cpu* CPU) {
	bit := GetNBit(cpu.A.Value, 7)
	cpu.A.Value <<= 1
	cpu.A.Value |= cpu.P.GetFlagC()
	updateFlagsROL(cpu, bit, cpu.A.Value)
}

func ExecuteROLMEM(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	bit := GetNBit(value, 7)
	value <<= 1
	value |= cpu.P.GetFlagC()
	updateFlagsROL(cpu, bit, value)
	cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, value)
	cpu.PC.Increment(1)
}

func ROLA(cpu* CPU) uint16 {
	ExecuteROLA(cpu)
	return 2
}

func ROLZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteROLMEM(cpu)
	return 5
}

func ROLZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteROLMEM(cpu)
	return 6
}

func ROLAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteROLMEM(cpu)
	return 6
}

func ROLAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteROLMEM(cpu)
	return 7
}

//ROR
func updateFlagsROR(cpu* CPU, bit byte, result byte) {
	cpu.P.SetFlagC(bit)
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
}

func ExecuteRORA(cpu* CPU) {
	bit := GetNBit(cpu.A.Value, 0)
	cpu.A.Value >>= 1
	cpu.A.Value |= (cpu.P.GetFlagC() << 7)
	updateFlagsROR(cpu, bit, cpu.A.Value)
}

func ExecuteRORMEM(cpu* CPU) {
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	bit := GetNBit(value, 0)
	value >>= 1
	value |= (cpu.P.GetFlagC() << 7)
	updateFlagsROR(cpu, bit, value)
	cpu.Mapper.Write(cpu.Addr.ADH, cpu.Addr.ADL, value)
	cpu.PC.Increment(1)
}

func RORA(cpu* CPU) uint16 {
	ExecuteRORA(cpu)
	return 2
}

func RORZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteRORMEM(cpu)
	return 5
}

func RORZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteRORMEM(cpu)
	return 6
}

func RORAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteRORMEM(cpu)
	return 6
}

func RORAbsoluteX(cpu* CPU) uint16 {
	GetAbsoluteXAddr(cpu)
  ExecuteRORMEM(cpu)
	return 7
}

//AND
func updateFlagsAND(cpu* CPU, result byte) {
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
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
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
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
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
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

//Boolean

//Compare
func updateFlagsCP(cpu* CPU, reg_val byte, mem_val byte, result byte) {
  if reg_val >= mem_val {
		cpu.P.SetFlagC(1)
	}else {
		cpu.P.SetFlagC(0)
	}
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(result, 7))
}

func ExecuteCP(cpu* CPU, reg* Register) {
  reg_val := reg.Value
	value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	result := reg_val - value
	updateFlagsCP(cpu, reg_val, value, result)
	cpu.PC.Increment(1)
}

//CMP
func CMPImmediate(cpu* CPU) uint16 {
  ExecuteCP(cpu, &cpu.A)
	return 2
}

func CMPZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 3
}

func CMPZeroPageX(cpu* CPU) uint16 {
	GetZeroPageXAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 4
}

func CMPAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 4
}

func CMPAbsoluteX(cpu* CPU) uint16 {
	adjusted := GetAbsoluteXAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 4 + adjusted
}

func CMPAbsoluteY(cpu* CPU) uint16 {
	adjusted := GetAbsoluteYAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 4 + adjusted
}

func CMPIndirectX(cpu* CPU) uint16 {
	GetIndirectXAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 6
}

func CMPIndirectY(cpu* CPU) uint16 {
	adjusted := GetIndirectYAddr(cpu)
  ExecuteCP(cpu, &cpu.A)
	return 5 + adjusted
}

//CPX
func CPXImmediate(cpu* CPU) uint16 {
  ExecuteCP(cpu, &cpu.X)
	return 2
}

func CPXZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteCP(cpu, &cpu.X)
	return 3
}

func CPXAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteCP(cpu, &cpu.X)
	return 4
}

//CPY
func CPYImmediate(cpu* CPU) uint16 {
  ExecuteCP(cpu, &cpu.Y)
	return 2
}

func CPYZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteCP(cpu, &cpu.Y)
	return 3
}

func CPYAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteCP(cpu, &cpu.Y)
	return 4
}
//Load Registers

func updateFlagsLoadRegister(cpu* CPU, value byte) {
	if value == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(value, 7))
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
	if value == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(value, 7))
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
	if value == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagN(GetNBit(value, 7))
}

func ExecutePull(cpu* CPU) byte {
  cpu.SP.Value++
  return cpu.Mapper.Read(0x01, cpu.SP.Value)
}

func ExecutePullA(cpu* CPU) {
  cpu.A.Value = ExecutePull(cpu)
  updateFlagsPull(cpu, cpu.A.Value)
}

func ExecutePullP(cpu* CPU) {
  cpu.P.Value = ExecutePull(cpu)
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

//Flag Operations

//CLC
func CLC(cpu* CPU) uint16 {
  cpu.P.SetFlagC(0)
  return 2
}

//SEC
func SEC(cpu* CPU) uint16 {
  cpu.P.SetFlagC(1)
  return 2
}

//CLD
func CLD(cpu* CPU) uint16 {
  cpu.P.SetFlagD(0)
  return 2
}

//SED
func SED(cpu* CPU) uint16 {
  cpu.P.SetFlagD(1)
  return 2
}

//CLI
func CLI(cpu* CPU) uint16 {
  cpu.P.SetFlagI(0)
  return 2
}

//SEI
func SEI(cpu* CPU) uint16 {
  cpu.P.SetFlagI(1)
  return 2
}

//CLV
func CLV(cpu* CPU) uint16 {
  cpu.P.SetFlagV(0)
  return 2
}

//Program Control

//Jump
func ExecuteJMP(cpu* CPU) {
	cpu.PC.ADL = cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
	cpu.Addr.Increment(1)
	cpu.PC.ADH = cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
}

func JMPAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
	ExecuteJMP(cpu)
	return 3
}

func JMPIndirect(cpu* CPU) uint16 {
	GetIndirectAddr(cpu)
	ExecuteJMP(cpu)
	return 5
}

//JSR
func ExecuteJSR(cpu* CPU) {
  aux_adl := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.PC.Increment(1)
  aux_adh := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  ExecutePush(cpu, cpu.PC.ADH)
  ExecutePush(cpu, cpu.PC.ADL)
  cpu.PC.ADL = aux_adl
  cpu.PC.ADH = aux_adh
}

func JSR(cpu* CPU) uint16 {
  ExecuteJSR(cpu)
  return 6
}

//RTS
func ExecuteRTS(cpu* CPU) {
  cpu.PC.ADL = ExecutePull(cpu)
  cpu.PC.ADH = ExecutePull(cpu)
  cpu.PC.Increment(1)
}

func RTS(cpu* CPU) uint16 {
  ExecuteRTS(cpu)
  return 6
}

//BRK
func ExecuteBRK(cpu* CPU) {
  cpu.PC.Increment(1)
  ExecutePush(cpu, cpu.PC.ADH)
  ExecutePush(cpu, cpu.PC.ADL)
  ExecutePush(cpu, (cpu.P.Value|0x10))
  cpu.PC.ADL = 0xFE
  cpu.PC.ADH = 0xFF
}

func BRK(cpu* CPU) uint16 {
  ExecuteBRK(cpu)
  return 6
}

//RTI
func ExecuteRTI(cpu* CPU) {
  cpu.PC.ADL = ExecutePull(cpu)
  cpu.PC.ADH = ExecutePull(cpu)
}

func RTI(cpu* CPU) uint16 {
  ExecuteRTI(cpu)
  return 6
}

//BIT
func updateFlagsBIT(cpu* CPU, result byte) {
	if result == 0x00 {
		cpu.P.SetFlagZ(1)
	}else {
		cpu.P.SetFlagZ(0)
	}
	cpu.P.SetFlagV(GetNBit(result, 6))
	cpu.P.SetFlagN(GetNBit(result, 7))
}

func ExecuteBIT(cpu* CPU) {
  value := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
  updateFlagsBIT(cpu, (cpu.A.Value & value))
  cpu.PC.Increment(1)
}

func BITZeroPage(cpu* CPU) uint16 {
  GetZeroPageAddr(cpu)
  ExecuteBIT(cpu)
	return 3
}

func BITAbsolute(cpu* CPU) uint16 {
	GetAbsoluteAddr(cpu)
  ExecuteBIT(cpu)
	return 4
}

//Branching
func ExecuteBranch(cpu* CPU, flag_cond bool) uint16 {
  offset := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  var adjusted, branched uint16 = 0, 0
  if flag_cond {
    branched = 1
    if GetNBit(offset, 7) == 1 {
      adjusted = cpu.PC.Decrement(-offset)
    }else {
      adjusted = cpu.PC.Increment(offset)
    }
  }
  cpu.PC.Increment(1)
  return branched+adjusted
}

//BCC
func BCC(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagC() == 0)
  return 2 + cycles
}

//BCS
func BCS(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagC() == 1)
  return 2 + cycles
}

//BEQ
func BEQ(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagZ() == 1)
  return 2 + cycles
}

//BMI
func BMI(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagN() == 1)
  return 2 + cycles
}

//BNE
func BNE(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagZ() == 0)
  return 2 + cycles
}

//BPL
func BPL(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagN() == 0)
  return 2 + cycles
}

//BVC
func BVC(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagV() == 0)
  return 2 + cycles
}

//BVS
func BVS(cpu* CPU) uint16 {
  cycles := ExecuteBranch(cpu, cpu.P.GetFlagV() == 1)
  return 2 + cycles
}

func (cpu* CPU) Execute(n_cycles uint16) {
	var current_cycles uint16 = 0
	for current_cycles < n_cycles {
		op_code := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
		cpu.PC.Increment(1)
		current_cycles += func_table[op_code](cpu)
	}
}
