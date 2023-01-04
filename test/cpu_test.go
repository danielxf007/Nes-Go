package test

import (
	"testing"
	"nes_go/processor"
	"nes_go/mappers"
)

func TestCPUANDImmediate(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var test_table = []struct {
		A processor.Register
		pc processor.ProgramCounter
		address uint16
		value byte
		result byte
		expected_flag processor.FlagRegister
	}{
		{processor.Register{0x80}, processor.ProgramCounter{0x01, 0xFF}, 0x01FF, 0xFF, 0x80, processor.FlagRegister{N: 1, Z: 0}},
		{processor.Register{0x80}, processor.ProgramCounter{0x02, 0x10}, 0x0210, 0x00, 0x00, processor.FlagRegister{N: 0, Z: 1}},
	}
	t.Log("Given the need to test the ANDImmediate operation.")
	{
		for _, element := range test_table {
			cpu.A = element.A
			cpu.PC = element.pc
			mapper.Write(cpu.PC.ADH, cpu.PC.ADL, element.value)
			cpu.P.Reset()
			processor.ANDImmediate(cpu)
			if cpu.A.Value == element.result {
				if cpu.P.N == element.expected_flag.N {
					if cpu.P.Z == element.expected_flag.Z {
						t.Log("\tOK")
					}else {
						t.Errorf("\tThere was a problem with the flag Z, got %d expected %d", cpu.P.Z, element.expected_flag.Z)
					}
				}else{
					t.Errorf("\tThere was a problem with the flag N, got %d expected %d", cpu.P.N, element.expected_flag.N)
				}
			}else {
				t.Errorf("\tThere was a problem with the result, got %d expected %d", cpu.A.Value, element.value)
			}
		}
	}
}

func TestCPUExecute(t *testing.T) {
	cpu := new(processor.CPU)
	t.Log("Given the need to test the execution loop.")
	{
		cpu.Execute(10)
	}
}