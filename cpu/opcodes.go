package cpu

type OpcodeFunc func(c *CPU)

func ADC(c *CPU) {}

// AND performs a AND of accumulator and data at memory location
func AND(c *CPU) {
	c.accumulator = c.accumulator & c.read(c.addr_absolute)
	if c.accumulator == 0x00 {
		c.setFlag(zero, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}
func ASL(c *CPU) {}
func BCC(c *CPU) {}
func BCS(c *CPU) {}
func BEQ(c *CPU) {}
func BIT(c *CPU) {}
func BMI(c *CPU) {}
func BNE(c *CPU) {}
func BPL(c *CPU) {}
func BRK(c *CPU) {}
func BVC(c *CPU) {}
func BVS(c *CPU) {}
func CLC(c *CPU) {}
func CLD(c *CPU) {}
func CLI(c *CPU) {}
func CLV(c *CPU) {}
func CMP(c *CPU) {}
func CPX(c *CPU) {}
func CPY(c *CPU) {}
func DEC(c *CPU) {}
func DEX(c *CPU) {}
func DEY(c *CPU) {}

// EOR performs an exclusive OR of accumulator and data at memory location
func EOR(c *CPU) {
	c.accumulator = c.accumulator ^ c.read(c.addr_absolute)
	if c.accumulator == 0x00 {
		c.setFlag(zero, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// INC adds one to the data at memory location
func INC(c *CPU) {
	data := c.read(c.addr_absolute) + 1
	c.write(c.addr_absolute, data)
	if data == 0x00 {
		c.setFlag(zero, true)
	}
	if data&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// INX adds one to the value in the X register
func INX(c *CPU) {
	c.regX = c.regX + 1
	if c.regX == 0x00 {
		c.setFlag(zero, true)
	}
	if c.regX&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// INX adds one to the value in the X register
func INY(c *CPU) {
	c.regY = c.regY + 1
	if c.regY == 0x00 {
		c.setFlag(zero, true)
	}
	if c.regY&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// JMP sets the program counter to the address at the operand
func JMP(c *CPU) {
}
func JSR(c *CPU) {}

// LDA loads the data at memory location into the accumulator
func LDA(c *CPU) {
	c.accumulator = c.read(c.addr_absolute)
	if c.accumulator == 0x00 {
		c.setFlag(zero, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// LDX loads the data at memory location into the X register
func LDX(c *CPU) {
	c.regX = c.read(c.addr_absolute)
	if c.regX == 0x00 {
		c.setFlag(zero, true)
	}
	if c.regX&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// LDY loads the data at memory location into the y register
func LDY(c *CPU) {
	c.regY = c.read(c.addr_absolute)
	if c.regY == 0x00 {
		c.setFlag(zero, true)
	}
	if c.regY&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}
func LSR(c *CPU) {}
func NOP(c *CPU) {}

// ORA performs an OR of the data at the accumulator and memory location
func ORA(c *CPU) {
	c.accumulator = c.accumulator | c.read(c.addr_absolute)
	if c.accumulator == 0x00 {
		c.setFlag(zero, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}
func PHA(c *CPU) {}
func PHP(c *CPU) {}
func PLA(c *CPU) {}
func PLP(c *CPU) {}
func ROL(c *CPU) {}
func ROR(c *CPU) {}
func RTI(c *CPU) {}
func RTS(c *CPU) {}
func SBC(c *CPU) {}

// SEC sets the carry flag to 1
func SEC(c *CPU) {
	c.setFlag(carry, true)
}

// SED sets the decimal flag to 1
func SED(c *CPU) {
	c.setFlag(decimal_mode, true)
}

// SEI sets the interrupt disable flag to 1
func SEI(c *CPU) {
	c.setFlag(irq_disable, true)
}

// STA stores the content of the accumulator into memory
func STA(c *CPU) {
	c.write(c.addr_absolute, c.accumulator)
}

// STX stores the content of the X registor into memory
func STX(c *CPU) {
	c.write(c.addr_absolute, c.regX)
}

// STY stores the content of the Y register into memory
func STY(c *CPU) {
	c.write(c.addr_absolute, c.regY)
}

// TAX copies the content of the X register to the accumulator
func TAX(c *CPU) {
	c.accumulator = c.regX
	if c.accumulator == 0x00 {
		c.setFlag(zero, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}

// TAY copies the content of the Y register to the accumulator
func TAY(c *CPU) {
	c.accumulator = c.regY
	if c.accumulator == 0x00 {
		c.setFlag(zero, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negative, true)
	}
}
func TSX(c *CPU) {}
func TXA(c *CPU) {}
func TXS(c *CPU) {}
func TYA(c *CPU) {}
