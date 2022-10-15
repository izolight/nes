package cpu

type AddressModeFunc func(c *CPU)

// Accumulator addressing
func Accum(c *CPU) {
	c.fetched = c.accumulator
}

// Immediate addressing
func Imm(c *CPU) {
	c.fetched = c.readPC()
}

// absolute addressing
func Abs(c *CPU) {
	lo := c.readPC()
	hi := c.readPC()
	c.addr_absolute = addr16(lo, hi)
}

// zero page addressing
func Zp(c *CPU) {
	c.addr_absolute = uint16(c.readPC())
}

// indexed zero page addressing x
func ZpX(c *CPU) {
	c.addr_absolute = uint16(c.readPC() + c.regX)
}

// indexed zero page addressing y
func ZpY(c *CPU) {
	c.addr_absolute = uint16(c.readPC() + c.regY)
}

// indexed absolute addressing x
func AbsX(c *CPU) {
	lo := c.readPC()
	hi := c.readPC()
	c.addr_absolute = addr16(lo, hi) + uint16(c.regX)
}

// indexed absolute addressing y
func AbsY(c *CPU) {
	lo := c.readPC()
	hi := c.readPC()
	c.addr_absolute = addr16(lo, hi) + uint16(c.regY)
}

// implied addressing
func Implied(c *CPU) {
}

// relative addressing
func Relative(c *CPU) {
	offset := c.readPC()
	c.addr_absolute = c.programCounter + uint16(offset)
}

// indexed indirect addressing x
func IndirectX(c *CPU) {
	ptr := uint16(c.readPC())
	lo := uint16(c.read((ptr + uint16(c.regX)) & 0x00FF))
	hi := uint16(c.read((ptr + uint16(c.regX) + 1) & 0x00FF))
	c.addr_absolute = hi<<8 | lo
}

// indirect indexed addressing y
func IndirectY(c *CPU) {
	ptr := uint16(c.readPC())
	lo := uint16(c.read(ptr & 0x00FF))
	hi := uint16(c.read((ptr + 1) & 0x00FF))
	c.addr_absolute = (hi<<8 | lo) + uint16(c.regY)
}

// absolut indrect addressing
func Indirect(c *CPU) {
	ptrLo := c.readPC()
	ptrHi := c.readPC()
	ptr := addr16(ptrLo, ptrHi)
	lo := c.read(ptr)
	hi := c.read(ptr + 1)
	c.addr_absolute = addr16(lo, hi)
}
