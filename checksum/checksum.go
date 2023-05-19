package checksum

func Calc(buf []byte) byte {
	var sum byte
	for _, b := range buf {
		sum += b
	}
	return sum
}
