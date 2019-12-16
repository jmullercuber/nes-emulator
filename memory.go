package main

type MemoryMap16Bit interface {
	write(address uint16, value uint8)
	read(address uint16) uint8
}

type CPUMemoryMap struct {
	// TODO: fill in
}
