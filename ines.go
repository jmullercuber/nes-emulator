package main

import (
	"fmt"
)

func ProcessINES(header ROMHeader, rom_bytes []byte) {
	const KB = 1024
	var prg_rom_size = uint(header[4]) //* 16 * KB
	var chr_rom_size = uint(header[5]) //* 8 * KB

	fmt.Println("PRG ROM Size:", prg_rom_size)
	fmt.Println("CHR ROM Size:", chr_rom_size)

	var has_trainer = (header[6] & 4) >> 2
	var trainer_size = uint(has_trainer) * 512
	fmt.Println("Trainer Available:", has_trainer != 0)
	fmt.Println("Trainer Size:", trainer_size)
	// flags 6
	// mirroring =
	// flags 7
	// flags 8
	// flags 9
	// flags 10
	// rest (11-15) usually empty padding

	// output should be configured memory map
}
