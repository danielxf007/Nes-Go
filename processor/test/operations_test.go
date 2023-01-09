package test

import (
	"testing"
	"nes_go/processor"
	"nes_go/mappers"
)

func TestExecuteAND(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
		A processor.Register
		value addrValue
		result byte
		flag processor.FlagRegister
	}{
	  {processor.Register{0x0F}, addrValue{0x01, 0x00, 0xFF}, 0x0F, processor.FlagRegister{N: 0, Z: 0}},
	  {processor.Register{0x80}, addrValue{0x02, 0xFF, 0xFF}, 0x80, processor.FlagRegister{N: 1, Z: 0}},
	  {processor.Register{0x80}, addrValue{0x03, 0xAA, 0x00}, 0x00, processor.FlagRegister{N: 0, Z: 1}},
	}
	t.Log("Given the need to test the ExecuteAND operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    processor.ExecuteAND(cpu)
	    t.Logf("Context A:%08b Mem[0x%02x%02x]:%08b", cpu.A.Value, value.ADH, value.ADL, value.Value)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result %08b", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got %08b expected %08b", cpu.A.Value, element.result)
	    }
	    if cpu.P.N == element.flag.N {
	      t.Logf("Got the expected N Flag %d", element.flag.N)
	    }else {
	      t.Errorf("There was a problem with the flag N, got %d expected %d", cpu.P.N, element.flag.N)
	    }
	    if cpu.P.Z == element.flag.Z {
	      t.Logf("Got the expected Z Flag %d", element.flag.Z)
	    }else {
	      t.Errorf("There was a problem with the flag Z, got %d expected %d", cpu.P.Z, element.flag.Z)
	    }
	  }
	}
}

func TestExecuteORA(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
		A processor.Register
		value addrValue
		result byte
		flag processor.FlagRegister
	}{
	  {processor.Register{0x1F}, addrValue{0x01, 0x00, 0x0F}, 0x1F, processor.FlagRegister{N: 0, Z: 0}},
	  {processor.Register{0x80}, addrValue{0x02, 0xFF, 0xFF}, 0xFF, processor.FlagRegister{N: 1, Z: 0}},
	  {processor.Register{0x00}, addrValue{0x03, 0xAA, 0x00}, 0x00, processor.FlagRegister{N: 0, Z: 1}},
	}
	t.Log("Given the need to test the ExecuteORA operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    processor.ExecuteORA(cpu)
	    t.Logf("Context A:%08b Mem[0x%02x%02x]:%08b", cpu.A.Value, value.ADH, value.ADL, value.Value)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result %08b", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got %08b expected %08b", cpu.A.Value, element.result)
	    }
	    if cpu.P.N == element.flag.N {
	      t.Logf("Got the expected N Flag %d", element.flag.N)
	    }else {
	      t.Errorf("There was a problem with the flag N, got %d expected %d", cpu.P.N, element.flag.N)
	    }
	    if cpu.P.Z == element.flag.Z {
	      t.Logf("Got the expected Z Flag %d", element.flag.Z)
	    }else {
	      t.Errorf("There was a problem with the flag Z, got %d expected %d", cpu.P.Z, element.flag.Z)
	    }
	  }
	}
}

func TestExecuteASLA(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var test_table = []struct {
		A processor.Register
		result byte
		flag processor.FlagRegister
	}{
	  {processor.Register{0x0F}, 0x1E, processor.FlagRegister{N: 0, Z: 0, C:0}},
	  {processor.Register{0x00}, 0x00, processor.FlagRegister{N: 0, Z: 1, C:0}},
	  {processor.Register{0x40}, 0x80, processor.FlagRegister{N: 1, Z: 0, C:0}},
	  {processor.Register{0x8F}, 0x1E, processor.FlagRegister{N: 0, Z: 0, C:1}},
	  {processor.Register{0x80}, 0x00, processor.FlagRegister{N: 0, Z: 1, C:1}},
	  {processor.Register{0xC1}, 0x82, processor.FlagRegister{N: 1, Z: 0, C:1}},
	}
	t.Log("Given the need to test the ExecuteASLA operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    processor.ASLA(cpu)
	    t.Logf("Context A:%08b", cpu.A.Value)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result %08b", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got %08b expected %08b", cpu.A.Value, element.result)
	    }
	    if cpu.P.N == element.flag.N {
	      t.Logf("Got the expected N Flag %d", element.flag.N)
	    }else {
	      t.Errorf("There was a problem with the flag N, got %d expected %d", cpu.P.N, element.flag.N)
	    }
	    if cpu.P.Z == element.flag.Z {
	      t.Logf("Got the expected Z Flag %d", element.flag.Z)
	    }else {
	      t.Errorf("There was a problem with the flag Z, got %d expected %d", cpu.P.Z, element.flag.Z)
	    }
	    if cpu.P.C == element.flag.C {
	      t.Logf("Got the expected C Flag %d", element.flag.C)
	    }else {
	      t.Errorf("There was a problem with the flag C, got %d expected %d", cpu.P.C, element.flag.C)
	    }
	  }
	}
}

func TestExecuteASLMEM(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var mem_value byte
	var test_table = []struct {
		value addrValue
		result byte
		flag processor.FlagRegister
	}{
	  {addrValue{0x01, 0x00, 0x0F}, 0x1E, processor.FlagRegister{N: 0, Z: 0, C:0}},
	  {addrValue{0x02, 0x00, 0x00}, 0x00, processor.FlagRegister{N: 0, Z: 1, C:0}},
	  {addrValue{0x03, 0x00, 0x40}, 0x80, processor.FlagRegister{N: 1, Z: 0, C:0}},
	  {addrValue{0x04, 0x00, 0x8F}, 0x1E, processor.FlagRegister{N: 0, Z: 0, C:1}},
	  {addrValue{0x05, 0x00, 0x80}, 0x00, processor.FlagRegister{N: 0, Z: 1, C:1}},
	  {addrValue{0x06, 0x00, 0xC1}, 0x82, processor.FlagRegister{N: 1, Z: 0, C:1}},
	}
	t.Log("Given the need to test the ExecuteASLA operation.")
	{
	  for _, element := range test_table {
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    processor.ExecuteASLMEM(cpu)
	    mem_value = mapper.Read(value.ADH, value.ADL)
	    t.Logf("Context Mem[0x%02x%02x]:%08b", value.ADH, value.ADL, value.Value)
	    if mem_value == element.result {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:%08b", value.ADH, value.ADL, mem_value)
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:%08b expected %08b",
	      value.ADH, value.ADL, mem_value, element.result)
	    }
	    if cpu.P.N == element.flag.N {
	      t.Logf("Got the expected N Flag %d", element.flag.N)
	    }else {
	      t.Errorf("There was a problem with the flag N, got %d expected %d", cpu.P.N, element.flag.N)
	    }
	    if cpu.P.Z == element.flag.Z {
	      t.Logf("Got the expected Z Flag %d", element.flag.Z)
	    }else {
	      t.Errorf("There was a problem with the flag Z, got %d expected %d", cpu.P.Z, element.flag.Z)
	    }
	    if cpu.P.C == element.flag.C {
	      t.Logf("Got the expected C Flag %d", element.flag.C)
	    }else {
	      t.Errorf("There was a problem with the flag C, got %d expected %d", cpu.P.C, element.flag.C)
	    }
	  }
	}
}

func TestExecuteLoadRegister(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	reg_types := [3]*processor.Register{&cpu.A, &cpu.X, &cpu.Y}
	reg_names := [3]string{"A", "X", "Y"}
	var test_table = []struct {
		value addrValue
		reg_type byte
		result byte
		flag processor.FlagRegister
	}{
	  {addrValue{0x01, 0x00, 0x0F}, 0, 0x0F, processor.FlagRegister{N: 0, Z: 0}},
	  {addrValue{0x02, 0x00, 0x00}, 1, 0x00, processor.FlagRegister{N: 0, Z: 1}},
	  {addrValue{0x03, 0x00, 0x80}, 2, 0x80, processor.FlagRegister{N: 1, Z: 0}},
	}
	t.Log("Given the need to test the ExecuteLoadRegister operation.")
	{
	  for _, element := range test_table {
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    processor.ExecuteLoadRegister(cpu, reg_types[element.reg_type])
	    t.Logf("Context Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, value.Value)
	    if reg_types[element.reg_type].Value == element.result {
	      t.Logf("Got the expected result %s:0x%02x", reg_names[element.reg_type], element.result)
	    }else {
	      t.Errorf("There was a problem with the result, %s:0x%02x expected %s:0x%02x",
	      reg_names[element.reg_type], reg_types[element.reg_type].Value, reg_names[element.reg_type], element.result)
	    }
	    if cpu.P.N == element.flag.N {
	      t.Logf("Got the expected N Flag %d", element.flag.N)
	    }else {
	      t.Errorf("There was a problem with the flag N, got %d expected %d", cpu.P.N, element.flag.N)
	    }
	    if cpu.P.Z == element.flag.Z {
	      t.Logf("Got the expected Z Flag %d", element.flag.Z)
	    }else {
	      t.Errorf("There was a problem with the flag Z, got %d expected %d", cpu.P.Z, element.flag.Z)
	    }
	  }
	}
}

func TestStoreRegister(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var addr processor.AddressBuffer
	reg_types := [3]*processor.Register{&cpu.A, &cpu.X, &cpu.Y}
	reg_names := [3]string{"A", "X", "Y"}
	var test_table = []struct {
	  reg_val byte
	  reg_type byte
	  addr processor.AddressBuffer
	}{
	  {0xFF, 0, processor.AddressBuffer{0x01, 0x00}},
	  {0xFA, 1, processor.AddressBuffer{0x01, 0x00}},
	  {0xFB, 2, processor.AddressBuffer{0x01, 0x00}},
	}
	t.Log("Given the need to test the ExecuteStoreRegister operation.")
	{
	  for _, element := range test_table {
	    reg_types[element.reg_type].Value = element.reg_val
	    addr = element.addr
	    cpu.Addr.ADH = addr.ADH
	    cpu.Addr.ADL = addr.ADL
	    processor.ExecuteStoreRegister(cpu, reg_types[element.reg_type])
	    t.Logf("Context %s:0x%02x Addr:0x%02x%02x", reg_names[element.reg_type], element.reg_val, addr.ADH, addr.ADL)
	    if reg_types[element.reg_type].Value == mapper.Read(addr.ADH, addr.ADL) {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:0x%02x", addr.ADH, addr.ADL, element.reg_val)
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:0x%02x expected Mem[0x%02x%02x]:0x%02x",
	      addr.ADH, addr.ADL, mapper.Read(addr.ADH, addr.ADL), addr.ADH, addr.ADL, element.reg_val)
	    }
	  }
	}
}
