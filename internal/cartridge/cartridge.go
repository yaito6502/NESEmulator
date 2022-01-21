package cartridge

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/yaito6502/NESEmulator/internal/mem"
)

type Cartridge struct {
	ProgramRom   mem.ROM
	CharacterRom mem.ROM
}

func NewCartridge(nespath string) *Cartridge {
	cart := new(Cartridge)
	cart.extractROMData(nespath)
	return cart
}

//[TODO] Headerの情報を16バイト全部使用して、各種設定を行う
func (cart *Cartridge) extractROMData(nespath string) {
	contents, err := ioutil.ReadFile(nespath)
	if err != nil {
		log.Fatal("Cannot read file")
	}
	header := contents[0:16]
	//check file format
	if filepath.Ext(nespath) != ".nes" || string(header[0:3]) != "NES" {
		log.Fatal("This file is NOT .nes format")
	}

	const NESHEADERSIZE int = 0x0010
	character_romstart := NESHEADERSIZE + 0x4000*int(header[4])
	character_romend := character_romstart + 0x2000*int(header[5])

	cart.ProgramRom = mem.NewROM(contents[NESHEADERSIZE : character_romstart-1])
	cart.CharacterRom = mem.NewROM(contents[character_romstart : character_romend-1])
}
