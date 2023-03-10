package processor

//Addressing modes

func GetZeroPageAddr(cpu* CPU) {
	cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
	cpu.Addr.ADH = 0x00
}

func GetZeroPageXAddr(cpu* CPU) {
  cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.Addr.ADL += cpu.X.Value
  cpu.Addr.ADH = 0x00
}

func GetZeroPageYAddr(cpu* CPU) {
  cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.Addr.ADL += cpu.Y.Value
  cpu.Addr.ADH = 0x00
}

func GetAbsoluteAddr(cpu* CPU) {
  cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.PC.Increment(1)
  cpu.Addr.ADH = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
}

func GetAbsoluteXAddr(cpu* CPU) uint16 {
  cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.PC.Increment(1)
  cpu.Addr.ADH = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  return cpu.Addr.Increment(cpu.X.Value)
}

func GetAbsoluteYAddr(cpu* CPU) uint16 {
  cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.PC.Increment(1)
  cpu.Addr.ADH = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  return cpu.Addr.Increment(cpu.Y.Value)
}

func GetIndirectXAddr(cpu* CPU) {
  aux_adl := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  aux_adl += cpu.X.Value
  cpu.Addr.ADL = cpu.Mapper.Read(0x00, aux_adl)
  aux_adl++
  cpu.Addr.ADH = cpu.Mapper.Read(0x00, aux_adl)
}

func GetIndirectYAddr(cpu* CPU) uint16 {
  aux_adl := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.Addr.ADL = cpu.Mapper.Read(0x00, aux_adl)
  aux_adl++
  cpu.Addr.ADH = cpu.Mapper.Read(0x00, aux_adl)
  return cpu.Addr.Increment(cpu.Y.Value)
}

func GetIndirectAddr(cpu* CPU) {
  cpu.Addr.ADL = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  cpu.PC.Increment(1)
  cpu.Addr.ADH = cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
  aux_adl := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
  cpu.Addr.Increment(1)
  aux_adh := cpu.Mapper.Read(cpu.Addr.ADH, cpu.Addr.ADL)
  cpu.Addr.ADH, cpu.Addr.ADL = aux_adh, aux_adl
}

//Bit
func GetNBit(value byte, n byte) byte {
  return ((value >> n) & 0x01) 
}

func SetNBit(value byte, n byte, bit byte) byte {
  return ((value & ^(0x01 << n)) | (bit << n)) 
}
