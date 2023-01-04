package processor

//Boolean

func updateFlagsAND(cpu* CPU) {
	cpu.P.N = (cpu.A.Value & 0x80) >> 7
	if cpu.A.Value == 0x00 {
		cpu.P.Z = 1
	}else {
		cpu.P.Z = 0
	}
}

func ANDImmediate(cpu* CPU) uint16 {
	value := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
	cpu.A.Value &= value
	updateFlagsAND(cpu)
	cpu.PC.Increment(1)
	return 2
}

func (cpu* CPU) Execute(n_cycles uint16) {
	var current_cycles uint16 = 0
	for current_cycles < n_cycles {
		op_code := cpu.Mapper.Read(cpu.PC.ADH, cpu.PC.ADL)
		cpu.PC.Increment(1)
		current_cycles += func_table[op_code](cpu)
	}
}