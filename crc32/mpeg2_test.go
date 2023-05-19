package crc32

import (
	"encoding/hex"
	"testing"
)

func TestMPEG2(t *testing.T) {
	var expected uint32 = 0x9EA45415
	hexStr := "1a2b3c"
	data, _ := hex.DecodeString(hexStr)
	result := MPEG2(data)
	if result != expected {
		t.Errorf("got %x, want %x", result, expected)
	}

	expected = 0x56D1FD12
	src := []byte("Hello, I'm Onur!")
	encodedStr := hex.EncodeToString(src) // Result = 48656c6c6f2c2049276d204f6e757221
	data, _ = hex.DecodeString(encodedStr)
	result = MPEG2(data)
	if result != expected {
		t.Errorf("got %x, want %x", result, expected)
	}

	expected = 0x9BA02E2D
	buf := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x01, 0x37, 0x38, 0x55}
	result = MPEG2(buf)
	if result != expected {
		t.Errorf("got %x, want %x", result, expected)
	}

	expected = 0xD3EC404D
	buf3 := []byte{0x55, 0xaa, 0x00, 0x16, 0x00, 0x0c, 0x35, 0x38, 0x43, 0x46, 0x37, 0x39, 0x30, 0x35, 0x46, 0x35, 0x42, 0x45, 0x05, 0xf5, 0xe1, 0x03, 0x04, 0xa6, 0xf7, 0xc1, 0xf2}
	result = MPEG2(buf3)
	if result != expected {
		t.Errorf("got %x, want %x", result, expected)
	}
}
