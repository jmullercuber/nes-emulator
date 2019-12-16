package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type ROMHeader [16]byte

var knownFileTypes = [2]string{"iNES", "NES 2.0"}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: nes-emulator <rom_path>")
		fmt.Println("  rom_path - Path to the game rom")
		os.Exit(1)
	}
	var input_rom_path = os.Args[1]
	fmt.Println("Reading ROM:", input_rom_path)

	var rom_bytes, err = loadBinaryFile(input_rom_path)
	if err != nil {
		panic("Unable to read ROM file")
	}

	fmt.Printf("% x\n", rom_bytes[:16])
	fmt.Printf("% x\n", rom_bytes[16:32])

	var header = getHeader(rom_bytes)
	var file_type = determineFileType(header)
	fmt.Println("File type:", file_type)
	fmt.Println()
	if file_type == "iNES" {
		ProcessINES(header, rom_bytes)
	} else {
		panic("Cannot process file type: " + file_type)
	}
}

func loadBinaryFile(file_path string) ([]byte, error) {
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

func getHeader(file_bytes []byte) ROMHeader {
	var header ROMHeader
	copy(header[:], file_bytes[:16])
	return header
}

func determineFileType(header ROMHeader) string {
	var iNESFormatStart = []byte{0x4e, 0x45, 0x53, 0x1a}
	if bytes.Equal(header[0:4], iNESFormatStart) {
		if (header[7] & 0x0C) == 0x08 {
			return "NES 2.0"
		} else {
			return "iNES"
		}
	}
	return "Unknown"
}
