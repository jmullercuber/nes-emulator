package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: nes-emulator <rom_path>")
		fmt.Println("  rom_path - Path to the game rom")
		os.Exit(1)
	}
	var input_rom_path = os.Args[1]
	fmt.Println("Reading ROM:", input_rom_path)

	var rom_bytes, err = LoadBinaryFile(input_rom_path)
	if err != nil {
		panic("Unable to read ROM file")
	}

	fmt.Println(rom_bytes[:16])
	fmt.Println(rom_bytes[16:32])
}

func LoadBinaryFile(file_path string) ([]byte, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var file_size int64 = stats.Size()
	file_bytes := make([]byte, file_size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(file_bytes)
	if err != nil {
		return nil, err
	}

	return file_bytes, nil
}
