package cpu

type OpcodeFunc func(c *CPU)

// ADC performs add with carry of the accumulator and data at memory location
func ADC(c *CPU) {
	result := uint16(c.accumulator) + uint16(c.read(c.addr_absolute)) + uint16(c.getFlag(carryFlag))
	carry := result > 255
	c.setFlag(carryFlag, carry)
	zero := (result & 0x00FF) == 0
	c.setFlag(zeroFlag, zero)
	// TODO overflow flag
	negative := (result & 0x80) == 1
	c.setFlag(negativeFlag, negative)
	c.accumulator = uint8(result & 0x00FF)
}

// AND performs a AND of the accumulator and data at memory location
func AND(c *CPU) {
	c.accumulator = c.accumulator & c.read(c.addr_absolute)
	c.setFlag(zeroFlag, c.accumulator == 0x00)
	c.setFlag(negativeFlag, c.accumulator&0x80 == 0x80)
}

// ASL performs an arithmetic left shift
func ASL(c *CPU) {
	data := c.read(c.addr_absolute)
	result := data << 1
	carry := result == 0 && data != 0
	c.setFlag(carryFlag, carry)
	c.setFlag(zeroFlag, result == 0)
	c.setFlag(negativeFlag, result&0x80 == 0x80)
	c.accumulator = result
}

// BCC branches if the carry flag is clear
func BCC(c *CPU) {
	if c.getFlag(carryFlag) == 0 {
		branch(c)
	}
}

// BCS branches if the carry flag is set
func BCS(c *CPU) {
	if c.getFlag(carryFlag) == 1 {
		branch(c)
	}
}

// BEQ branches if the zero flag is set
func BEQ(c *CPU) {
	if c.getFlag(zeroFlag) == 1 {
		branch(c)
	}
}
func BIT(c *CPU) {}

// BMI branches if the negative flag is set
func BMI(c *CPU) {
	if c.getFlag(negativeFlag) == 1 {
		branch(c)
	}
}

// BNE branches if the zero flag is clear
func BNE(c *CPU) {
	if c.getFlag(zeroFlag) == 0 {
		branch(c)
	}
}

// BPL branches if the negative flag is clear
func BPL(c *CPU) {
	if c.getFlag(negativeFlag) == 0 {
		branch(c)
	}
}

// branch is a helper for the branch instructions
func branch(c *CPU) {
	c.remainingCycles++
	// going to new page, additional cycle required
	if (c.addr_absolute & 0xFF00) != (c.programCounter & 0xFF00) {
		c.remainingCycles++
	}
	c.programCounter = c.addr_absolute
}

func BRK(c *CPU) {}

// BVC branches if the overflow flag is clear
func BVC(c *CPU) {
	if c.getFlag(overflowFlag) == 0 {
		branch(c)
	}
}

// BVS branches if the overflow flag is set
func BVS(c *CPU) {
	if c.getFlag(overflowFlag) == 1 {
		branch(c)
	}
}

// CLC clears the carry flag
func CLC(c *CPU) {
	c.setFlag(carryFlag, false)
}

// CLD clears the decimal flag
func CLD(c *CPU) {
	c.setFlag(decimalModeFlag, false)
}

// CLI clears the interrupt disableflag
func CLI(c *CPU) {
	c.setFlag(irqDisableFlag, false)
}

// CLV clears the overflow flag
func CLV(c *CPU) {
	c.setFlag(overflowFlag, false)
}

// CMP compares the accumulator with the data at the memory location
func CMP(c *CPU) {
	result := c.accumulator - c.read(c.addr_absolute)
	c.setFlag(zeroFlag, result == 0x00)
	// TODO: rest
}
func CPX(c *CPU) {}
func CPY(c *CPU) {}

// DEC substracts one from the data at the memory location
func DEC(c *CPU) {
	data := c.read(c.addr_absolute) - 1
	c.setFlag(zeroFlag, data == 0x00)
	c.setFlag(negativeFlag, data&0x80 == 0x80)
	c.write(c.addr_absolute, data)
}

// DEX decrements the value in the X register
func DEX(c *CPU) {
	c.regX = c.regX - 1
	c.setFlag(zeroFlag, c.regX == 0x00)
	c.setFlag(negativeFlag, c.regX&0x80 == 0x80)
}

// DEY decrements the value in the Y register
func DEY(c *CPU) {
	c.regY = c.regY - 1
	c.setFlag(zeroFlag, c.regY == 0x00)
	c.setFlag(negativeFlag, c.regY&0x80 == 0x80)
}

// EOR performs an exclusive OR of accumulator and data at memory location
func EOR(c *CPU) {
	c.accumulator = c.accumulator ^ c.read(c.addr_absolute)
	c.setFlag(zeroFlag, c.accumulator == 0x00)
	c.setFlag(negativeFlag, c.accumulator&0x80 == 0x80)
}

// INC adds one to the data at memory location
func INC(c *CPU) {
	data := c.read(c.addr_absolute) + 1
	c.write(c.addr_absolute, data)
	if data == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if data&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// INX adds one to the value in the X register
func INX(c *CPU) {
	c.regX = c.regX + 1
	if c.regX == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.regX&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// INX adds one to the value in the X register
func INY(c *CPU) {
	c.regY = c.regY + 1
	if c.regY == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.regY&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// JMP sets the program counter to the address at the operand
func JMP(c *CPU) {
	c.programCounter = c.addr_absolute
}

// JSR pushes the pc(-1) on the stack and sets the pc to the new address,
func JSR(c *CPU) {
	c.programCounter--
	hi := uint8((c.programCounter >> 8) & 0x00FF)
	lo := uint8(c.programCounter & 0x00FF)
	c.push(hi)
	c.push(lo)
	c.programCounter = c.addr_absolute
}

// LDA loads the data at memory location into the accumulator
func LDA(c *CPU) {
	c.accumulator = c.read(c.addr_absolute)
	if c.accumulator == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// LDX loads the data at memory location into the X register
func LDX(c *CPU) {
	c.regX = c.read(c.addr_absolute)
	if c.regX == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.regX&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// LDY loads the data at memory location into the y register
func LDY(c *CPU) {
	c.regY = c.read(c.addr_absolute)
	if c.regY == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.regY&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}
func LSR(c *CPU) {}

// NOP does nothing (just incrementing pc)
func NOP(c *CPU) {}

// ORA performs an OR of the data at the accumulator and memory location
func ORA(c *CPU) {
	c.accumulator = c.accumulator | c.read(c.addr_absolute)
	if c.accumulator == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// PHA pushes a copy of the accumulator on the stack
func PHA(c *CPU) {
	c.push(c.accumulator)
}

// PHP pushes a copy of the status flags to the stack
func PHP(c *CPU) {
	c.push(c.statusReg)
}

// PLA pulls the stack value into the accumulator
func PLA(c *CPU) {
	c.accumulator = c.pull()
	c.setFlag(zeroFlag, c.accumulator == 0x00)
	c.setFlag(negativeFlag, c.accumulator&0x80 == 0x80)
}

// PLP pulls the stack value into the status flags
func PLP(c *CPU) {
	c.statusReg = c.pull()
}
func ROL(c *CPU) {}
func ROR(c *CPU) {}
func RTI(c *CPU) {}

// RTS pulls the program counter(-1) from th stack
func RTS(c *CPU) {
	lo := c.pull()
	hi := c.pull()
	c.programCounter = addr16(lo, hi)
	c.programCounter++
}
func SBC(c *CPU) {}

// SEC sets the carry flag to 1
func SEC(c *CPU) {
	c.setFlag(carryFlag, true)
}

// SED sets the decimal flag to 1
func SED(c *CPU) {
	c.setFlag(decimalModeFlag, true)
}

// SEI sets the interrupt disable flag to 1
func SEI(c *CPU) {
	c.setFlag(irqDisableFlag, true)
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
		c.setFlag(zeroFlag, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}

// TAY copies the content of the Y register to the accumulator
func TAY(c *CPU) {
	c.accumulator = c.regY
	if c.accumulator == 0x00 {
		c.setFlag(zeroFlag, true)
	}
	if c.accumulator&0x80 == 0x80 {
		c.setFlag(negativeFlag, true)
	}
}
func TSX(c *CPU) {}
func TXA(c *CPU) {}
func TXS(c *CPU) {}
func TYA(c *CPU) {}
