package lrc

func Calc(buf []byte) byte {
	var lrc byte
	for _, b := range buf {
		lrc += b
	}
	return (lrc ^ 0xff) + 1
}
