package main

type ROMHeader struct {
	entry          [4]uint8
	logo           [0x30]uint8
	title          [16]byte
	newLicCode     uint16
	sgbFlag        uint8
	romType        uint8
	romSize        uint8
	ramSize        uint8
	destCode       uint8
	licCode        uint8
	version        uint8
	checksum       uint8
	globalChecksum uint16
}

type Cartridge struct {
	char      []byte
	romSize   uint32
	romData   *uint8
	romHeader *ROMHeader
}

func (cart *Cartridge) loadCartridge() {
	print()
}
