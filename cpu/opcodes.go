package cpu

type OpcodeFunc func(c *CPU)

func ADC(c *CPU) {}
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
func EOR(c *CPU) {}
func INC(c *CPU) {}
func INX(c *CPU) {}
func INY(c *CPU) {}
func JMP(c *CPU) {}
func JSR(c *CPU) {}
func LDA(c *CPU) {}
func LDX(c *CPU) {}
func LDY(c *CPU) {}
func NOP(c *CPU) {}
func ORA(c *CPU) {}
func PHA(c *CPU) {}
func PHP(c *CPU) {}
func PLA(c *CPU) {}
func PLP(c *CPU) {}
func ROL(c *CPU) {}
func ROR(c *CPU) {}
func RTI(c *CPU) {}
func RTS(c *CPU) {}
func SBC(c *CPU) {}
func SEC(c *CPU) {}
func SED(c *CPU) {}
func SEI(c *CPU) {}
func STA(c *CPU) {}
func STX(c *CPU) {}
func STY(c *CPU) {}
func TAX(c *CPU) {}
func TAY(c *CPU) {}
func TSX(c *CPU) {}
func TXA(c *CPU) {}
func TXS(c *CPU) {}
func TYA(c *CPU) {}
