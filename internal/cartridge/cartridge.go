package cartridge

import (
	"io/ioutil"
	"log"
)

/*
type Cartridge struct {
	program_rom   []byte
	character_rom []byte
}*/

func AttachCartridge(filename string) ([]byte, []byte) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal()
	}

	const NESHEADERSIZE int = 0x0010

	character_romstart := NESHEADERSIZE + 0x4000*int(bytes[4])
	character_romend := character_romstart + 0x2000*int(bytes[5])

	program_rom := bytes[NESHEADERSIZE : character_romstart-1]
	character_rom := bytes[character_romstart : character_romend-1]
	return program_rom, character_rom
}
