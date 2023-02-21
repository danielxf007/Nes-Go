package memory

import (
	"errors"
	"os"
)

type RomHeader struct {
	PGR byte
	CHR byte
	Flags_6 byte
	Flags_7 byte
	PGR_ram_sz byte
	Flags_9 byte
	Flags_10 byte
}

type RomOperations interface {
	hasMinimunSz(data []byte) bool
	isNESFile(file_type []byte) bool
	hasUnusedBytes(unused_bytes []byte) bool
	Load(rom_path string) error
	HeaderToString() string
	GetPGR() byte
	GetCHR() byte
	GetFlags6() (byte, byte, byte, byte)
	GetFlags7() (byte, byte, byte, byte)
	GetPGRRamSz() byte
	GetFlags9() (byte, byte)
	GetFlags10() (byte, byte, byte)
}

type Rom struct {
	Header RomHeader
	PGR_banks [][]byte
	CHR_banks [][]byte
}

//Check if the file has at least a header
func (rom* Rom) hasMinimunSz(data []byte) bool {
	return len(data) >= 16
}

//Check if the file has the first 4 characters that identify a nes file
func (rom* Rom) isNESFile(file_type []byte) bool {
	return file_type[0] == 'N' && file_type[1] == 'E' && file_type[2] == 'S' && file_type[3] == 0x1A
}

//Check the bytes 8 to 15 are 0x00, programs do not function otherwise
func (rom* Rom) hasUnusedBytes(unused_bytes []byte) bool {
	for index := range unused_bytes {
		if unused_bytes[index] != 0x00 {
			return false
		}
	}
	return true
}

//Read and load the content of a rom file 
func (rom* Rom) Load(rom_path string) error {
	data, err := os.ReadFile(rom_path)
	if err != nil {
		return errors.New("The file could not be loaded")
	}
	if !rom.hasMinimunSz(data) {
		return errors.New("The file must be at least 16 bytes long")
	}
	if !rom.isNESFile(data[0:4]) {
		return errors.New("This is not a NES file")
	}
	if !rom.hasUnusedBytes(data[8:16]) {
		return errors.New("Bytes from 8 to 15 must be 0x00")
	}
	rom.Header = RomHeader {
		PGR: data[4], CHR: data[5], Flags_6: data[6],
		Flags_7: data[7], PGR_ram_sz: data[8], Flags_9: data[9], Flags_10: data[10],
	}
	var start, end, offset, pgr_unit_sz, chr_unit_sz uint32
	pgr_unit_sz = 0x4000
	rom.PGR_banks = make([][]byte, rom.Header.PGR)
	for index := range rom.PGR_banks {
	  rom.PGR_banks[index] = make([]byte, pgr_unit_sz)
	}
	offset = 16
	for index := range rom.PGR_banks {
	    start = (uint32(index)*pgr_unit_sz) + offset
	    end = start + pgr_unit_sz
	    rom.PGR_banks[index] = data[start:end]
	}
	chr_unit_sz = 0x2000
	rom.CHR_banks = make([][]byte, rom.Header.CHR)
	for index := range rom.CHR_banks {
	  rom.CHR_banks[index] = make([]byte, chr_unit_sz)
	}
	offset = (pgr_unit_sz*uint32(rom.Header.PGR)) + 16
	for index := range rom.CHR_banks {
	    start = (uint32(index)*chr_unit_sz) + offset
	    end = start + chr_unit_sz
	    rom.CHR_banks[index] = data[start:end]
	}
	return nil
}

func (rom* Rom) Read(address uint16) byte {
	return 0x00
}

func (rom* Rom) Write(address uint16, value byte) {
}

func (rom* Rom) Reset() {
}




