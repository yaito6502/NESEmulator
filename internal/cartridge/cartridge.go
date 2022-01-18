package cartridge

import (
	"io/ioutil"
	"log"
)

type Cartridge struct {
	ProgramRom   []byte
	CharacterRom []byte
}

const NESHEADERSIZE int = 0x0010

func NewCartridge(filename string) *Cartridge {
	cart := new(Cartridge)
	cart.ProgramRom, cart.CharacterRom = attachCartridge(filename)
	return cart
}

func attachCartridge(filename string) ([]byte, []byte) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal()
	}

	//check nes format
	if string(contents[0:3]) != "NES" {
		log.Fatal()
	}

	character_romstart := NESHEADERSIZE + 0x4000*int(contents[4])
	character_romend := character_romstart + 0x2000*int(contents[5])

	program_rom := contents[NESHEADERSIZE : character_romstart-1]
	character_rom := contents[character_romstart : character_romend-1]
	return program_rom, character_rom
}
