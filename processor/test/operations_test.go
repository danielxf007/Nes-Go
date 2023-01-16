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
		flags byte
	}{
	  {processor.Register{0x0F}, addrValue{0x01, 0x00, 0xFF}, 0x0F, 0b00000000},
	  {processor.Register{0x80}, addrValue{0x02, 0xFF, 0xFF}, 0x80, 0b10000000},
	  {processor.Register{0x80}, addrValue{0x03, 0xAA, 0x00}, 0x00, 0b00000010},
	}
	t.Log("Given the need to test the ExecuteAND operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    cpu.P.Reset()
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
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
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
		flags byte
	}{
	  {processor.Register{0x1F}, addrValue{0x01, 0x00, 0x0F}, 0x1F, 0b00000000},
	  {processor.Register{0x80}, addrValue{0x02, 0xFF, 0xFF}, 0xFF, 0b10000000},
	  {processor.Register{0x00}, addrValue{0x03, 0xAA, 0x00}, 0x00, 0b00000010},
	}
	t.Log("Given the need to test the ExecuteORA operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    cpu.P.Reset()
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
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
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
		flags byte
	}{
	  {processor.Register{0x0F}, 0x1E, 0b00000000},
	  {processor.Register{0x00}, 0x00, 0b00000010},
	  {processor.Register{0x40}, 0x80, 0b10000000},
	  {processor.Register{0x8F}, 0x1E, 0b00000001},
	  {processor.Register{0x80}, 0x00, 0b00000011},
	  {processor.Register{0xC1}, 0x82, 0b10000001},
	}
	t.Log("Given the need to test the ExecuteASLA operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    cpu.P.Reset()
	    processor.ASLA(cpu)
	    t.Logf("Context A:%08b", cpu.A.Value)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result %08b", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got %08b expected %08b", cpu.A.Value, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
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
		flags byte
	}{
	  {addrValue{0x01, 0x00, 0x0F}, 0x1E, 0b00000000},
	  {addrValue{0x02, 0x00, 0x00}, 0x00, 0b00000010},
	  {addrValue{0x03, 0x00, 0x40}, 0x80, 0b10000000},
	  {addrValue{0x04, 0x00, 0x8F}, 0x1E, 0b00000001},
	  {addrValue{0x05, 0x00, 0x80}, 0x00, 0b00000011},
	  {addrValue{0x06, 0x00, 0xC1}, 0x82, 0b10000001},
	}
	t.Log("Given the need to test the ExecuteASLMEM operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context Mem[0x%02x%02x]:%08b", value.ADH, value.ADL, value.Value)
	    processor.ExecuteASLMEM(cpu)
	    mem_value = mapper.Read(value.ADH, value.ADL)
	    if mem_value == element.result {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:%08b", value.ADH, value.ADL, mem_value)
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:%08b expected %08b",
	      value.ADH, value.ADL, mem_value, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecuteLSRA(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var test_table = []struct {
		A processor.Register
		result byte
		flags byte
	}{
	  {processor.Register{0x02}, 0x01, 0b00000000},
	  {processor.Register{0x00}, 0x00, 0b00000010},
	  {processor.Register{0x81}, 0x40, 0b00000001},
	  {processor.Register{0x01}, 0x00, 0b00000011},
	}
	t.Log("Given the need to test the ExecuteLSRA operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    cpu.P.Reset()
	    processor.ExecuteLSRA(cpu)
	    t.Logf("Context A:%08b", cpu.A.Value)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result %08b", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got %08b expected %08b", cpu.A.Value, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecuteLSRMEM(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var mem_value byte
	var test_table = []struct {
		value addrValue
		result byte
		flags byte
	}{
	  {addrValue{0x01, 0x00, 0x02}, 0x01, 0b00000000},
	  {addrValue{0x02, 0x00, 0x00}, 0x00, 0b00000010},
	  {addrValue{0x04, 0x00, 0x81}, 0x40, 0b00000001},
	  {addrValue{0x05, 0x00, 0x01}, 0x00, 0b00000011},
	}
	t.Log("Given the need to test the ExecuteLSRMEM operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context Mem[0x%02x%02x]:%08b", value.ADH, value.ADL, value.Value)
	    processor.ExecuteLSRMEM(cpu)
	    mem_value = mapper.Read(value.ADH, value.ADL)
	    if mem_value == element.result {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:%08b", value.ADH, value.ADL, mem_value)
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:%08b expected %08b",
	      value.ADH, value.ADL, mem_value, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
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
		flags byte
	}{
	  {addrValue{0x01, 0x00, 0x0F}, 0, 0x0F, 0b00000000},
	  {addrValue{0x02, 0x00, 0x00}, 1, 0x00, 0b00000010},
	  {addrValue{0x03, 0x00, 0x80}, 2, 0x80, 0b10000000},
	}
	t.Log("Given the need to test the ExecuteLoadRegister operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, value.Value)
	    processor.ExecuteLoadRegister(cpu, reg_types[element.reg_type])
	    if reg_types[element.reg_type].Value == element.result {
	      t.Logf("Got the expected result %s:0x%02x", reg_names[element.reg_type], element.result)
	    }else {
	      t.Errorf("There was a problem with the result, %s:0x%02x expected %s:0x%02x",
	      reg_names[element.reg_type], reg_types[element.reg_type].Value, reg_names[element.reg_type], element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecuteStoreRegister(t *testing.T) {
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
	    t.Logf("Context %s:0x%02x Addr:0x%02x%02x", reg_names[element.reg_type], element.reg_val, addr.ADH, addr.ADL)
	    processor.ExecuteStoreRegister(cpu, reg_types[element.reg_type])
	    if reg_types[element.reg_type].Value == mapper.Read(addr.ADH, addr.ADL) {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:0x%02x", addr.ADH, addr.ADL, element.reg_val)
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:0x%02x expected Mem[0x%02x%02x]:0x%02x",
	      addr.ADH, addr.ADL, mapper.Read(addr.ADH, addr.ADL), addr.ADH, addr.ADL, element.reg_val)
	    }
	  }
	}
}

func TestExecuteTransferValueRegisters(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	reg_types := [3]*processor.Register{&cpu.A, &cpu.X, &cpu.Y}
	reg_names := [3]string{"A", "X", "Y"}
	var test_table = []struct {
	  from byte
	  to byte
	  value byte
	  flags byte
	}{
	  {0, 1, 0x0F, 0b00000000},
	  {0, 2, 0x00, 0b00000010},
	  {0, 1, 0x80, 0b10000000},
	  {1, 0, 0x0F, 0b00000000},
	  {1, 2, 0x00, 0b00000010},
	  {1, 0, 0x80, 0b10000000},
	  {2, 0, 0x0F, 0b00000000},
	  {2, 1, 0x00, 0b00000010},
	  {2, 0, 0x80, 0b10000000},
	}
	t.Log("Given the need to test the ExecuteTransferValueRegisters operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    reg_types[element.from].Value = element.value
	    t.Logf("Context %s:0x%02x %s:0x%02x", reg_names[element.from], reg_types[element.from].Value,
	    reg_names[element.to], reg_types[element.to].Value)
	    processor.ExecuteTransferValueRegisters(cpu, reg_types[element.from], reg_types[element.to])
	    if reg_types[element.from].Value == reg_types[element.to].Value {
	      t.Logf("Got the expected result %s:0x%02x", reg_names[element.to], reg_types[element.to].Value)
	    }else {
	      t.Errorf("There was a problem with the result, got %s:0x%02x expected %s:0x%02x",
	      reg_names[element.to], reg_types[element.to].Value, reg_names[element.to], reg_types[element.from].Value)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecutePush(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
	  value byte
	  sp_value_i byte
	  sp_value_f byte
	  result addrValue
	}{
	  {0x0F, 0x5F, 0x5E, addrValue{0x01, 0x5F, 0x0F}},
	  {0x0F, 0x00, 0xFF, addrValue{0x01, 0x00, 0x0F}},
	}
	t.Log("Given the need to test the ExecutePush operation.")
	{
	  for _, element := range test_table {
	    cpu.SP.Value = element.sp_value_i
	    value = element.result
	    t.Logf("Context SP:0x%02x Pushed_Value:0x%02x", cpu.SP.Value, element.value)
	    processor.ExecutePush(cpu, element.value)
	    if element.value == mapper.Read(value.ADH, value.ADL) {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, value.Value)
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:0x%02x expected Mem[0x%02x%02x]:0x%02x",
	      value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL), value.ADH, value.ADL, value.Value)
	    }
	    if cpu.SP.Value == element.sp_value_f {
	      t.Logf("Got the expected result SP:0x%02x", cpu.SP.Value)
	    }else {
	      t.Errorf("There was a problem with the result, got SP:0x%02x expected SP:0x%02x",
	      cpu.SP.Value, element.sp_value_f)
	    }
	  }
	}
}

func TestExecutePullA(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
	  value addrValue
	  sp_value_i byte
	  sp_value_f byte
	  flags byte
	}{
	  {addrValue{0x01, 0x60, 0x0F}, 0x5F, 0x60, 0b00000000},
	  {addrValue{0x01, 0x01, 0x00}, 0x00, 0x01, 0b00000010},
	  {addrValue{0x01, 0x00, 0x80}, 0xFF, 0x00, 0b10000000},
	}
	t.Log("Given the need to test the ExecutePullA operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.SP.Value = element.sp_value_i
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context SP:0x%02x Mem[0x%02x%02x]:0x%02x", cpu.SP.Value, value.ADH, value.ADL, value.Value)
	    processor.ExecutePullA(cpu)
	    if cpu.A.Value == value.Value {
	      t.Logf("Got the expected result %s:0x%02x", "A", value.Value)
	    }else {
	      t.Errorf("There was a problem with the result, got %s:0x%02x expected %s:0x%02x",
	      "A", cpu.A.Value, "A", value.Value)
	    }
	    if cpu.SP.Value == element.sp_value_f {
	      t.Logf("Got the expected result SP:0x%02x", cpu.SP.Value)
	    }else {
	      t.Errorf("There was a problem with the result, got SP:0x%02x expected SP:0x%02x",
	      cpu.SP.Value, element.sp_value_f)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecutePullP(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
	  value addrValue
	  sp_value_i byte
	  sp_value_f byte
	  flags byte
	}{
	  {addrValue{0x01, 0x60, 0xFF}, 0x5F, 0x60, 0b11011111},
	}
	t.Log("Given the need to test the ExecutePullA operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.SP.Value = element.sp_value_i
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context SP:0x%02x Mem[0x%02x%02x]:0x%02x", cpu.SP.Value, value.ADH, value.ADL, value.Value)
	    processor.ExecutePullP(cpu)
	    if cpu.P.Value == value.Value {
	      t.Logf("Got the expected result %s:0x%02x", "P", value.Value)
	    }else {
	      t.Errorf("There was a problem with the result, got %s:0x%02x expected %s:0x%02x",
	      "P", cpu.P.Value, "P", value.Value)
	    }
	    if cpu.SP.Value == element.sp_value_f {
	      t.Logf("Got the expected result SP:0x%02x", cpu.SP.Value)
	    }else {
	      t.Errorf("There was a problem with the result, got SP:0x%02x expected SP:0x%02x",
	      cpu.SP.Value, element.sp_value_f)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecuteEOR(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
		A processor.Register
		value addrValue
		result byte
		flags byte
	}{
	  {processor.Register{0x10}, addrValue{0x01, 0x00, 0x0F}, 0x1F, 0b00000000},
	  {processor.Register{0x8F}, addrValue{0x02, 0xFF, 0x0F}, 0x80, 0b10000000},
	  {processor.Register{0xFF}, addrValue{0x03, 0xAA, 0xFF}, 0x00, 0b00000010},
	}
	t.Log("Given the need to test the ExecuteEOR operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    cpu.P.Reset()
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context A:%08b Mem[0x%02x%02x]:%08b", cpu.A.Value, value.ADH, value.ADL, value.Value)
	    processor.ExecuteEOR(cpu)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result %08b", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got %08b expected %08b", cpu.A.Value, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecuteINCReg(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	reg_types := [2]*processor.Register{&cpu.X, &cpu.Y}
	reg_names := [2]string{"X", "Y"}
	var test_table = []struct {
		reg_val byte
		reg_type byte
		result byte
		flags byte
	}{
	  {0x0F, 0, 0x10, 0b00000000},
	  {0xFF, 1, 0x00, 0b00000010},
	  {0x7F, 0, 0x80, 0b10000000},
	}
	t.Log("Given the need to test the ExecuteINCReg operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    reg_types[element.reg_type].Value = element.reg_val
	    t.Logf("Context %s:0x%02x", reg_names[element.reg_type], reg_types[element.reg_type].Value)
	    processor.ExecuteINCReg(cpu, reg_types[element.reg_type])
	    if reg_types[element.reg_type].Value == element.result {
	      t.Logf("Got the expected result %s:0x%02x", reg_names[element.reg_type], element.result)
	    }else {
	      t.Errorf("There was a problem with the result, %s:0x%02x expected %s:0x%02x",
	      reg_names[element.reg_type], reg_types[element.reg_type].Value, reg_names[element.reg_type], element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}


func TestExecuteINCMEM(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
		value addrValue
		result byte
		flags byte
	}{
	  {addrValue{0x01, 0x60, 0x0F}, 0x10, 0b00000000},
	  {addrValue{0x01, 0x60, 0xFF}, 0x00, 0b00000010},
	  {addrValue{0x01, 0x60, 0x7F}, 0x80, 0b10000000},
	}
	t.Log("Given the need to test the ExecuteINCMEM operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL))
	    processor.ExecuteINCMEM(cpu)
	    if mapper.Read(value.ADH, value.ADL) == element.result {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL))
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:0x%02x expected Mem[0x%02x%02x]:0x%02x",
	      value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL), value.ADH, value.ADL, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}


func TestExecuteDECReg(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	reg_types := [2]*processor.Register{&cpu.X, &cpu.Y}
	reg_names := [2]string{"X", "Y"}
	var test_table = []struct {
		reg_val byte
		reg_type byte
		result byte
		flags byte
	}{
	  {0x10, 0, 0x0F, 0b00000000},
	  {0x01, 1, 0x00, 0b00000010},
	  {0x00, 0, 0xFF, 0b10000000},
	}
	t.Log("Given the need to test the ExecuteDECReg operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    reg_types[element.reg_type].Value = element.reg_val
	    t.Logf("Context %s:0x%02x", reg_names[element.reg_type], reg_types[element.reg_type].Value)
	    processor.ExecuteDECReg(cpu, reg_types[element.reg_type])
	    if reg_types[element.reg_type].Value == element.result {
	      t.Logf("Got the expected result %s:0x%02x", reg_names[element.reg_type], element.result)
	    }else {
	      t.Errorf("There was a problem with the result, %s:0x%02x expected %s:0x%02x",
	      reg_names[element.reg_type], reg_types[element.reg_type].Value, reg_names[element.reg_type], element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}


func TestExecuteDECMEM(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
		value addrValue
		result byte
		flags byte
	}{
	  {addrValue{0x01, 0x60, 0x10}, 0x0F, 0b00000000},
	  {addrValue{0x01, 0x60, 0x01}, 0x00, 0b00000010},
	  {addrValue{0x01, 0x60, 0x00}, 0xFF, 0b10000000},
	}
	t.Log("Given the need to test the ExecuteDECMEM operation.")
	{
	  for _, element := range test_table {
	    cpu.P.Reset()
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL))
	    processor.ExecuteDECMEM(cpu)
	    if mapper.Read(value.ADH, value.ADL) == element.result {
	      t.Logf("Got the expected result Mem[0x%02x%02x]:0x%02x", value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL))
	    }else {
	      t.Errorf("There was a problem with the result, got Mem[0x%02x%02x]:0x%02x expected Mem[0x%02x%02x]:0x%02x",
	      value.ADH, value.ADL, mapper.Read(value.ADH, value.ADL), value.ADH, value.ADL, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

func TestExecuteSBC(t *testing.T) {
  mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var value addrValue
	var test_table = []struct {
		A processor.Register
		value addrValue
		result byte
		flags byte
	}{
		{processor.Register{0x05}, addrValue{0x01, 0x00, 0x05}, 0x00, 0b00000011},//substract the same number 
	  {processor.Register{0x05}, addrValue{0x01, 0x00, 0x03}, 0x02, 0b00000001},//substract with borrow (c==1->no borrow); positive result 
	  {processor.Register{0x05}, addrValue{0x01, 0x00, 0x06}, 0xFF, 0b11000000},//substract with borrow (c==0->borrow); negative result
	}
	t.Log("Given the need to test the ExecuteSBC operation.")
	{
	  for _, element := range test_table {
	    cpu.A = element.A
	    cpu.P.Reset()
	    cpu.P.SetFlagC(1)
	    value = element.value
	    cpu.Addr.ADH = value.ADH
	    cpu.Addr.ADL = value.ADL
	    mapper.Write(value.ADH, value.ADL, value.Value)
	    t.Logf("Context A:0x%02x Mem[0x%02x%02x]:0x%02x", cpu.A.Value, value.ADH, value.ADL, value.Value)
	    processor.ExecuteSBC(cpu)
	    if cpu.A.Value == element.result {
	      t.Logf("Got the expected result 0x%02x", element.result)
	    }else {
	      t.Errorf("There was a problem with the result, got 0x%02x expected 0x%02x", cpu.A.Value, element.result)
	    }
	    if cpu.P.Value == element.flags {
	      t.Logf("Got the expected Flags %08b", cpu.P.Value)
	    }else {
	      t.Errorf("There was a problem with the flags, got %08b expected %08b", cpu.P.Value, element.flags)
	    }
	  }
	}
}

