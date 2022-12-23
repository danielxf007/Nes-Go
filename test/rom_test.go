package test

import (
	"testing"
	"nes_go/memory"
)

//TestLoadRom checks the loading operation for rom memory
func TestLoadRom(t *testing.T) {
	rom := new(memory.Rom)
	file_path := "../roms/Excitebike.nes"
	t.Log("Given the need to test loading binary data to rom.")
	{
		err := rom.Load(file_path)
		if err != nil {
			t.Log(err)
		}else {
			t.Log("The file was loaded")
		}
	}
}