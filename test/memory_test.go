package test

import (
	"testing"
	"math/rand"
	"time"
	"nes_go/memory"
)

const check_mark = "\u2713"
const ballot_x = "\u2717"

//TestReadLinearMemory checks the reading operation for linear memory
func TestReadLinearMemory(t *testing.T) {
	linear_memory := new(memory.LinearMemory)
	linear_memory.Data = make([]uint8, 256, 256)
	linear_memory.Data[0x00] = 180
	linear_memory.Data[0xF0] = 204
	linear_memory.Data[0xFF] = 10
	var test_table = []struct {
		address uint8
		value uint8
	}{
		{0x00, 180},
		{0xF0, 204},
		{0xFF, 10},
	}
	t.Log("Given the need to test reading a linear memory.")
	{
		for _, element := range test_table {
			t.Logf("\tChecking address %d", element.address)
			{
				if linear_memory.Data[element.address] != element.value {
					t.Errorf("\t\tThe address %d should have %d, intead it got %d %v", element.address, element.value, linear_memory.Data[element.address], ballot_x)
				}else {
					t.Logf("\t\tThe memory address %d had the correct value %d %v", element.address, linear_memory.Data[element.address], check_mark)
				}
			}
		}
	}
}

//TestReadLinearMemory checks the writting operation for linear memory
func TestWriteLinearMemory(t *testing.T) {
	linear_memory := new(memory.LinearMemory)
	linear_memory.Data = make([]uint8, 256, 256)
	linear_memory.Write(0x00, 180)
	linear_memory.Write(0xF0, 204)
	linear_memory.Write(0xFF, 10)
	var test_table = []struct {
		address uint8
		value uint8
	}{
		{0x00, 180},
		{0xF0, 204},
		{0xFF, 10},
	}
	t.Log("Given the need to test writting on linear memory.")
	{
		for _, element := range test_table {
			t.Logf("\tChecking address %d", element.address)
			{
				if linear_memory.Data[element.address] != element.value {
					t.Errorf("\t\tThe address %d should have %d, intead it got %d %v", element.address, element.value, linear_memory.Data[element.address], ballot_x)
				}else {
					t.Logf("\t\tThe memory address %d had the correct value %d %v", element.address, linear_memory.Data[element.address], check_mark)
				}
			}
		}
	}
}

//TestResetLinearMemory checks if all bytes are reseted
func TestResetLinearMemory(t *testing.T) {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	linear_memory := new(memory.LinearMemory)
	linear_memory.Data = make([]uint8, 256, 256)
	for index := range linear_memory.Data {
		linear_memory.Data[index] = uint8(r1.Intn(0xFF))
	}
	linear_memory.Reset()
	t.Log("Given the need to test reset on linear memory.")
	{
		for index := range linear_memory.Data {
			{
				if linear_memory.Data[index] != 0 {
					t.Fatalf("\t\tThe address %d should have 0, intead it got %d", index, linear_memory.Data[index])
				}
			}
		}
		t.Log("\t\tThe memory was cleared")
	}
}