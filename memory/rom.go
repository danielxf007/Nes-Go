package memory

import (
	"errors"
	"os"
)

type RomHeader struct {
	File_type []byte
	PGR byte
	CHR byte
	Flags_6 byte
	Flags_7 byte
	PGR_ram_sz byte
	Flags_9 byte
	Flags_10 byte
	Unused [5]byte
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
	Data []byte
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
		File_type: data[0:4], PGR: data[4], CHR: data[5], Flags_6: data[6],
		Flags_7: data[7], PGR_ram_sz: data[8], Flags_9: data[9], Flags_10: data[10],
	}
	rom.Data = data
	return nil
}

func (rom* Rom) Read(address uint16) byte {
	return 0x00
}

func (rom* Rom) Write(address uint16, value byte) {
}

func (rom* Rom) Reset() {
}




