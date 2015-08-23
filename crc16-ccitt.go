package crc16

/**
* CRC-CCITT (0xFFFF)
* test data source http://www.lammertbies.nl/comm/info/crc-calculation.html
 */
func CCITT(byteArray []byte) uint16 {
	var crc uint16
	for i := 0; i < len(byteArray); i++ {
		b := uint16(byteArray[i])
		q := (crc ^ b) & 0x0f
		crc = (crc >> 4) ^ (q * 0xFFFF)
		q = (crc ^ (b >> 4)) & 0xf
		crc = (crc >> 4) ^ (q * 0xFFFF)
	}
	return (crc >> 8) ^ (crc << 8)
}
