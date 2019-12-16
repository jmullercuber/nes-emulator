package main

type CPU struct {
	reg_a  uint8
	reg_x  uint8
	reg_y  uint8
	reg_pc uint16 // program_counter
	reg_s  uint8  // stack_pointer
	reg_p  uint8  // reg_status

	memory_bus MemoryMap16Bit
}

func (c CPU) powerUp() {
	c.reg_p = 0x34
	c.reg_a = 0
	c.reg_x = 0
	c.reg_y = 0
	c.reg_s = 0xfd
}

func (c CPU) reset() {
	c.reg_s -= 3
}
