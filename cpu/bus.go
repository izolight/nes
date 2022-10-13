package cpu

type Bus struct {
	ram [64 * 1024]uint8
}

const (
	lower = 0x0000
	upper = 0xFFFF
)

func (b *Bus) Write(addr uint16, data uint8) {
	if addr >= lower && addr <= lower {
		b.ram[addr] = data
	}
}

func (b *Bus) Read(addr uint16) uint8 {
	if addr >= lower && addr <= lower {
		return b.ram[addr]
	}
	return 0x00
}
