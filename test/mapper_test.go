package test

import (
	"testing"
	"nes_go/mappers"
)

func TestNoMapperRead(t *testing.T) {
	mapper := new(mappers.NoMapper)
	var test_table = []struct {
		page_number uint8
		offset uint8
		value uint8
	}{
		{0x00, 0x00, 0},
		{0x0A, 0x10, 0},
		{0x0B, 0x30, 0},
	}
	t.Log("Given the need to test reading using no mapper.")
	{
		for _, element := range test_table {
			t.Logf("\tChecking page %d with offset %d", element.page_number, element.offset)
			{
				if mapper.Read(element.page_number, element.offset) != element.value {
					t.Errorf("\t\tThe page number %d offset %d should have %d, instead it got %d", element.page_number, element.offset, element.value, mapper.Read(element.page_number, element.offset))
				}else {
					t.Logf("\t\tThe page number %d offset %d has the correct value %d", element.page_number, element.offset, element.value)
				}
			}
		}
	}
}