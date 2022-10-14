package cpu

// Accumulator addressing
func Accum(c *CPU) {
	c.fetched = c.accumulator
}

// Immediate addressing
func Imm(c *CPU) {
	c.fetched = c.read(c.programCounter)
	c.programCounter++
}

// absolute addressing
func Abs(c *CPU) {
	lo := uint16(c.read(c.programCounter))
	c.programCounter++
	hi := uint16(c.read(c.programCounter))
	c.programCounter++
	c.addr_absolute = hi<<8 | lo
}

// zero page addressing
func Zp(c *CPU) {
	addr := c.read(c.programCounter)
	c.programCounter++
	c.addr_absolute = uint16(addr)
}

// indexed zero page addressing x
func ZpX(c *CPU) {
	addr := c.read(c.programCounter) + c.regX
	c.programCounter++
	c.addr_absolute = uint16(addr)
}

// indexed zero page addressing y
func ZpY(c *CPU) {
	addr := c.read(c.programCounter) + c.regY
	c.programCounter++
	c.addr_absolute = uint16(addr)
}

// indexed absolute addressing x
func AbsX(c *CPU) {
	lo := uint16(c.read(c.programCounter))
	c.programCounter++
	hi := uint16(c.read(c.programCounter))
	c.programCounter++
	addr := hi<<8 | lo
	c.addr_absolute = addr + uint16(c.regX)
}

// indexed absolute addressing y
func AbsY(c *CPU) {
	lo := uint16(c.read(c.programCounter))
	c.programCounter++
	hi := uint16(c.read(c.programCounter))
	c.programCounter++
	addr := hi<<8 | lo
	c.addr_absolute = addr + uint16(c.regY)
}

// implied addressing
func Implied(c *CPU) {
}

// relative addressing
func Relative(c *CPU) {
	offset := uint16(c.read(c.programCounter))
	c.programCounter++
	c.addr_absolute = c.programCounter + offset
}

// indexed indirect addressing x
func IndirectX(c *CPU) {
	addr := uint16(c.read(c.programCounter))
	c.programCounter++
	lo := uint16(c.read((addr + uint16(c.regX)) & 0x00FF))
	hi := uint16(c.read((addr + uint16(c.regX) + 1) & 0x00FF))
	c.addr_absolute = hi<<8 | lo
}

// indirect indexed addressing y
func IndirectY(c *CPU) {
	addr := uint16(c.read(c.programCounter))
	c.programCounter++
	lo := uint16(c.read(addr & 0x00FF))
	hi := uint16(c.read((addr + 1) & 0x00FF))
	c.addr_absolute = (hi<<8 | lo) + uint16(c.regY)
}

// absolut indrect addressing
func Indirect(c *CPU) {
	plo := uint16(c.read(c.programCounter))
	c.programCounter++
	phi := uint16(c.read(c.programCounter))
	c.programCounter++
	addr := phi<<8 | plo
	lo := uint16(c.read(addr))
	hi := uint16(c.read(addr + 1))
	c.addr_absolute = hi<<8 | lo
}

func absoluteIndexedIndirect(c *CPU) {
	lo := uint16(c.read(c.programCounter))
	c.programCounter++
	hi := uint16(c.read(c.programCounter))
	c.programCounter++
	c.addr_absolute = (hi<<8 | lo) + uint16(c.regX)
}
