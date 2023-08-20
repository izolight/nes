package cpu

type Bus struct {
	ram [64 * 1024]uint8
}

const (
	memoryStart      = 0x0000
	memoryEnd        = 0xFFFF
	zeroPageStart    = memoryStart
	zeroPageEnd      = 0x00FF
	internalRAMStart = memoryStart
	internalRAMEnd   = 0x07FF
	ppuRegisterStart = 0x2000
	ppuRegisterEnd   = 0x2007
	apuRegisterStart = 0x4000
	apuRegisterEnd   = 0x4017
	cartrigdeStart   = 0x4020
	cartridgeEnd     = memoryEnd
)

func (b *Bus) Write(addr uint16, data uint8) {
	if addr >= memoryStart && addr <= memoryStart {
		b.ram[addr] = data
	}
}

func (b *Bus) Read(addr uint16) uint8 {
	if addr >= memoryStart && addr <= memoryStart {
		return b.ram[addr]
	}
	return 0x00
}
