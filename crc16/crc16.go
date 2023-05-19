package crc16

import "sync"

var crcTable []uint16
var mux sync.Mutex

func Modbus(data []byte) (crc uint16) {
	if crcTable == nil {
		// Thread safe initialization.
		mux.Lock()
		if crcTable == nil {
			initTable()
		}
		mux.Unlock()
	}

	crc = 0xffff
	for _, v := range data {
		crc = (crc >> 8) ^ crcTable[(crc^uint16(v))&0x00FF]
	}

	return crc
}

func initTable() {
	crc16IBM := uint16(0xA001)
	crcTable = make([]uint16, 256)

	for i := uint16(0); i < 256; i++ {

		crc := uint16(0)
		c := i

		for j := uint16(0); j < 8; j++ {
			if ((crc ^ c) & 0x0001) > 0 {
				crc = (crc >> 1) ^ crc16IBM
			} else {
				crc = crc >> 1
			}
			c = c >> 1
		}
		crcTable[i] = crc
	}
}
