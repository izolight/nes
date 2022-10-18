package cpu

type CPU struct {
	registers
	bus *Bus

	fetched         uint8
	instructionReg  uint8
	addr_absolute   uint16
	remainingCycles uint
	opcodes         [64]instruction
}

type instruction struct {
	name      string
	cycles    uint
	addrMode  AddressModeFunc
	operation OpcodeFunc
}

type registers struct {
	accumulator    uint8
	regX           uint8
	regY           uint8
	programCounter uint16
	stackPointer   uint8
	statusReg      uint8
}

func (c *CPU) getFlag(f flag) uint8 {
	if c.statusReg&uint8(f) > 0 {
		return 1
	}
	return 0
}

func (c *CPU) setFlag(f flag, val bool) {
	if val {
		c.statusReg |= uint8(f)
	} else {
		c.statusReg &= uint8(f)
	}
}

func (c *CPU) read(addr uint16) uint8 {
	return c.bus.Read(addr)
}

// readPC is a convenience function for reading from the program counter.
func (c *CPU) readPC() uint8 {
	addr := c.bus.Read(c.programCounter)
	c.programCounter++
	return addr
}

func (c *CPU) write(addr uint16, data uint8) {
	c.bus.Write(addr, data)
}

type flag uint8

const (
	carryFlag flag = 1 << iota
	zeroFlag
	irqDisableFlag
	decimalModeFlag
	brkCommandFlag
	unusedFlag
	overflowFlag
	negativeFlag
)

func (c *CPU) Clock() {
	if c.remainingCycles == 0 {
		c.instructionReg = c.read(c.programCounter)
		c.programCounter++
		instruction := c.opcodes[c.instructionReg]
		instruction.addrMode(c)
		instruction.operation(c)
		c.remainingCycles += instruction.cycles
	}
	c.remainingCycles--
}

// addr16 constructs the uint16 addr from the low and high bytes.
func addr16(lo uint8, hi uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}
