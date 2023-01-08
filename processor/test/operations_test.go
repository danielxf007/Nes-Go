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
