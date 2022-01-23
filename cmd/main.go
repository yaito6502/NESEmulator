package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/nes"
)

func main() {
	//[TODO]flagを使用し、debugモードやoriginal sizeの設定を可能にする
	if len(os.Args) == 1 {
		log.Fatal("Input File Not Found")
	}
	//tools.SpriteDump(os.Args[1])
	cart := cartridge.NewCartridge(os.Args[1])
	nes := nes.NewNES(cart)
	Test(nes)
	//nes.Run()
}

func Test(nes *nes.NES) {
	file, err := os.Open("../test/nestest.log")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	steps := 1
	for fileScanner.Scan() {
		nes.CPU.Run()
		fmt.Println(fileScanner.Text())
		nes.Info.Print()
		infos := strings.Split(fileScanner.Text(), " ")
		for i := 0; i < len(infos); i++ {
			if strings.HasPrefix(infos[i], "A:") {
				exp, _ := strconv.ParseUint(infos[i][2:], 16, 8)
				if exp != uint64(nes.Info.A) {
					log.Fatalf("Error at Step %d Reg A\nexp[%X] your[%X]\n", steps, exp, nes.Info.A)
				}
			}
			if strings.HasPrefix(infos[i], "X:") {
				exp, _ := strconv.ParseUint(infos[i][2:], 16, 8)
				if exp != uint64(nes.Info.X) {
					log.Fatalf("Error at Step %d Reg X\nexp[%X] your[%X]\n", steps, exp, nes.Info.X)
				}
			}
			if strings.HasPrefix(infos[i], "Y:") {
				exp, _ := strconv.ParseUint(infos[i][2:], 16, 8)
				if exp != uint64(nes.Info.Y) {
					log.Fatalf("Error at Step %d Reg Y\nexp[%X] your[%X]\n", steps, exp, nes.Info.Y)
				}
			}
			if strings.HasPrefix(infos[i], "P:") {
				exp, _ := strconv.ParseUint(infos[i][2:], 16, 8)
				if exp != uint64(nes.Info.P) {
					log.Fatalf("Error at Step %d Reg P\nexp[%X] your[%X]\n", steps, exp, nes.Info.P)
				}
			}
			if strings.HasPrefix(infos[i], "SP:") {
				exp, _ := strconv.ParseUint(infos[i][3:], 16, 8)
				if exp != uint64(nes.Info.SP) {
					log.Fatalf("Error at Step %d Reg SP\nexp[%X] your[%X]\n", steps, exp, nes.Info.SP)
				}
			}
		}
		//fmt.Println(fileScanner.Text())
		//nes.Info.Print()
		steps++
	}
	if nes.CPUBUS.Read(0x0002) != 0 || nes.CPUBUS.Read(0x0003) != 0 {
		log.Fatal("Error at Last Step")
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
