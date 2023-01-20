package test

import (
	"testing"
	"nes_go/processor"
	"nes_go/mappers"
)

func TestGetZeroPageAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2 addrValue
	var expected_addr processor.AddressBuffer
	var test_table = []struct {
	  pc processor.ProgramCounter
	  b2 addrValue
	  expected_addr processor.AddressBuffer
	}{
		{processor.ProgramCounter{0x01, 0xFF}, addrValue{0x01, 0xFF, 0x0A}, processor.AddressBuffer{0x00, 0x0A}},
	}
	t.Log("Given the need to test the Zero Page.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    b2 = element.b2
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    processor.GetZeroPageAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x Mem[0x%02x%02x]:0x%02x", cpu.PC.ADH, cpu.PC.ADL, b2.ADH, b2.ADL, b2.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	  }
	}
}

func TestGetZeroPageXAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2 addrValue
	var expected_addr processor.AddressBuffer
	var test_table = []struct {
	  pc processor.ProgramCounter
	  x processor.Register
	  b2 addrValue
	  expected_addr processor.AddressBuffer
	}{
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x0A}, addrValue{0x01, 0xFF, 0x00}, processor.AddressBuffer{0x00, 0x0A}},
	}
	t.Log("Given the need to test the Zero Page with X index.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    cpu.X = element.x
	    b2 = element.b2
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    processor.GetZeroPageXAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x X:0x%02x Mem[0x%02x%02x]:0x%02x", cpu.PC.ADH, cpu.PC.ADL, cpu.X.Value, b2.ADH, b2.ADL, b2.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	  }
	}
}

func TestGetAbsoluteAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2, b3 addrValue
	var expected_addr processor.AddressBuffer
	var test_table = []struct {
	  pc processor.ProgramCounter
	  b2 addrValue
	  b3 addrValue
	  expected_addr processor.AddressBuffer
	}{
		{processor.ProgramCounter{0x01, 0xFF}, addrValue{0x01, 0xFF, 0xF0}, addrValue{0x02, 0x00, 0xA0}, 
		processor.AddressBuffer{0xA0, 0xF0}},
	}
	t.Log("Given the need to test the Absolute Address.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    b2 = element.b2
	    b3 = element.b3
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    mapper.Write(b3.ADH, b3.ADL, b3.Value)
	    processor.GetAbsoluteAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x",
	    element.pc.ADH, element.pc.ADL, b2.ADH, b2.ADL, b2.Value, b3.ADH, b3.ADL, b3.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	  }
	}
}

func TestGetAbsoluteXAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2, b3 addrValue
	var expected_addr processor.AddressBuffer
	var carry uint16
	var test_table = []struct {
	  pc processor.ProgramCounter
	  x processor.Register
	  b2 addrValue
	  b3 addrValue
	  expected_addr processor.AddressBuffer
	  carry uint16
	}{
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x1F}, addrValue{0x01, 0xFF, 0xF0}, addrValue{0x02, 0x00, 0xA0}, 
		processor.AddressBuffer{0xA1, 0x0F}, 1},
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x0F}, addrValue{0x01, 0xFF, 0xF0}, addrValue{0x02, 0x00, 0xA0}, 
		processor.AddressBuffer{0xA0, 0xFF}, 0},
	}
	t.Log("Given the need to test the Absolute Address with X index.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    cpu.X = element.x
	    b2 = element.b2
	    b3 = element.b3
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    mapper.Write(b3.ADH, b3.ADL, b3.Value)
	    carry = processor.GetAbsoluteXAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x X:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x",
	    element.pc.ADH, element.pc.ADL, cpu.X.Value, b2.ADH, b2.ADL, b2.Value, b3.ADH, b3.ADL, b3.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	    if carry == element.carry {
	      t.Logf("Got the expected result Carry %d", carry)
	    }else {
	      t.Errorf("There was a problem with the carry, got %d expected %d", 
	      carry, element.carry)
	    }
	  }
	}
}

func TestGetAbsoluteYAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2, b3 addrValue
	var expected_addr processor.AddressBuffer
	var carry uint16
	var test_table = []struct {
	  pc processor.ProgramCounter
	  Y processor.Register
	  b2 addrValue
	  b3 addrValue
	  expected_addr processor.AddressBuffer
	  carry uint16
	}{
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x1F}, addrValue{0x01, 0xFF, 0xF0}, addrValue{0x02, 0x00, 0xA0}, 
		processor.AddressBuffer{0xA1, 0x0F}, 1},
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x0F}, addrValue{0x01, 0xFF, 0xF0}, addrValue{0x02, 0x00, 0xA0}, 
		processor.AddressBuffer{0xA0, 0xFF}, 0},
	}
	t.Log("Given the need to test the Absolute Address with Y index.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    cpu.Y = element.Y
	    b2 = element.b2
	    b3 = element.b3
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    mapper.Write(b3.ADH, b3.ADL, b3.Value)
	    carry = processor.GetAbsoluteYAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x Y:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x",
	    element.pc.ADH, element.pc.ADL, cpu.Y.Value, b2.ADH, b2.ADL, b2.Value, b3.ADH, b3.ADL, b3.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	    if carry == element.carry {
	      t.Logf("Got the expected result Carry %d", carry)
	    }else {
	      t.Errorf("There was a problem with the carry, got %d expected %d", 
	      carry, element.carry)
	    }
	  }
	}
}

func TestGetIndirectXAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2 addrValue
	var addr_adl, addr_adh addrValue
	var expected_addr processor.AddressBuffer
	var test_table = []struct {
	  pc processor.ProgramCounter
	  x processor.Register
	  b2 addrValue
	  addr_adl addrValue
	  addr_adh addrValue
	  expected_addr processor.AddressBuffer
	}{
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x01}, addrValue{0x01, 0xFF, 0x0E}, addrValue{0x00, 0x0F, 0xF0},
		addrValue{0x00, 0x10, 0xA0}, processor.AddressBuffer{0xA0, 0xF0}},
	}
	t.Log("Given the need to test the Indirect Address with X index.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    cpu.X = element.x
	    b2 = element.b2
	    addr_adl = element.addr_adl
	    addr_adh = element.addr_adh
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    mapper.Write(addr_adl.ADH, addr_adl.ADL, addr_adl.Value)
	    mapper.Write(addr_adh.ADH, addr_adh.ADL, addr_adh.Value)
	    processor.GetIndirectXAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x X:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x",
	    element.pc.ADH, element.pc.ADL, cpu.X.Value, b2.ADH, b2.ADL, b2.Value, addr_adl.ADH, addr_adl.ADL, addr_adl.Value,
	    addr_adh.ADH, addr_adh.ADL, addr_adh.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	  }
	}
}

func TestGetIndirectYAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2 addrValue
	var addr_adl, addr_adh addrValue
	var expected_addr processor.AddressBuffer
	var carry uint16
	var test_table = []struct {
	  pc processor.ProgramCounter
	  Y processor.Register
	  b2 addrValue
	  addr_adl addrValue
	  addr_adh addrValue
	  expected_addr processor.AddressBuffer
	  carry uint16
	}{
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x10}, addrValue{0x01, 0xFF, 0x0F}, addrValue{0x00, 0x0F, 0xF0},
		addrValue{0x00, 0x10, 0xA0}, processor.AddressBuffer{0xA1, 0x00}, 1},
		{processor.ProgramCounter{0x01, 0xFF}, processor.Register{0x01}, addrValue{0x01, 0xFF, 0x0F}, addrValue{0x00, 0x0F, 0x0F},
		addrValue{0x00, 0x10, 0xA0}, processor.AddressBuffer{0xA0, 0x10}, 0},
	}
	t.Log("Given the need to test the Indirect Address with Y index.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    cpu.Y = element.Y
	    b2 = element.b2
	    addr_adl, addr_adh = element.addr_adl, element.addr_adh
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    mapper.Write(addr_adl.ADH, addr_adl.ADL, addr_adl.Value)
	    mapper.Write(addr_adh.ADH, addr_adh.ADL, addr_adh.Value)
	    carry = processor.GetIndirectYAddr(cpu)
	    t.Logf("Context PC:0x%02x%02x Y:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x",
	    element.pc.ADH, element.pc.ADL, cpu.Y.Value, b2.ADH, b2.ADL, b2.Value, addr_adl.ADH, addr_adl.ADL, addr_adl.Value,
	    addr_adh.ADH, addr_adh.ADL, addr_adh.Value)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	    if carry == element.carry {
	      t.Logf("Got the expected result Carry %d", carry)
	    }else {
	      t.Errorf("There was a problem with the carry, got %d expected %d", 
	      carry, element.carry)
	    }
	  }
	}
}

func TestGetIndirectAddr(t *testing.T) {
	mapper := new(mappers.NoMapper)
	cpu := new(processor.CPU)
	cpu.Mapper = mapper
	var b2, b3, addr_l, addr_h addrValue
	var expected_addr processor.AddressBuffer
	var test_table = []struct {
	  pc processor.ProgramCounter
	  b2 addrValue
	  b3 addrValue
	  addr_l addrValue
	  addr_h addrValue
	  expected_addr processor.AddressBuffer
	}{
		{processor.ProgramCounter{0x01, 0xFF}, addrValue{0x01, 0xFF, 0x40}, addrValue{0x02, 0x00, 0x0A}, 
		addrValue{0x0A, 0x40, 0xFF}, addrValue{0x0A, 0x41, 0x03}, processor.AddressBuffer{0x03, 0xFF}},
	}
	t.Log("Given the need to test the Absolute Address.")
	{
	  for _, element := range test_table {
	    cpu.PC = element.pc
	    b2 = element.b2
	    b3 = element.b3
	    addr_l = element.addr_l
	    addr_h = element.addr_h
	    expected_addr = element.expected_addr
	    mapper.Write(b2.ADH, b2.ADL, b2.Value)
	    mapper.Write(b3.ADH, b3.ADL, b3.Value)
	    mapper.Write(addr_l.ADH, addr_l.ADL, addr_l.Value)
	    mapper.Write(addr_h.ADH, addr_h.ADL, addr_h.Value)
	    t.Logf("Context PC:0x%02x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x Mem[0x%02x%02x]:0x%02x",
	    cpu.PC.ADH, cpu.PC.ADL, b2.ADH, b2.ADL, b2.Value, b3.ADH, b3.ADL, b3.Value,
	    addr_l.ADH, addr_l.ADL, addr_l.Value, addr_h.ADH, addr_h.ADL, addr_h.Value)
	    processor.GetIndirectAddr(cpu)
	    if cpu.Addr.ADH == expected_addr.ADH && cpu.Addr.ADL == expected_addr.ADL {
	      t.Logf("Got the expected result Addres 0x%02x%02x", expected_addr.ADH, expected_addr.ADL)
	    }else {
	      t.Errorf("There was a problem with the Address, got %02x%02x expected %02x%02x", 
	      cpu.Addr.ADH, cpu.Addr.ADL, expected_addr.ADH, expected_addr.ADL)
	    }
	  }
	}
}
