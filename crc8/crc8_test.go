package crc8

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "TestCalc1",
			args: args{
				buf: []byte{0x01},
			},
			want: 0x07,
		},
		{
			name: "TestCalc2",
			args: args{
				buf: []byte{0x02},
			},
			want: 0x0E,
		},
		{
			name: "TestCalc2",
			args: args{
				buf: []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39},
			},
			want: 0xF4,
		},
		{
			name: "TestCalc3",
			args: args{
				buf: []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x01, 0x37, 0x38, 0x55},
			},
			want: 0x3C,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.buf); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
