package lrc

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
				buf: []byte{0x36, 0x37, 0x41, 0x46},
			},
			want: 0x0C,
		},
		{
			name: "TestCalc2",
			args: args{
				buf: []byte{0x01},
			},
			want: 0xFF,
		},
		{
			name: "TestCalc3",
			args: args{
				buf: []byte{0x36, 0x55, 0x41, 0x46},
			},
			want: 0xEE,
		},
		{
			name: "TestCalc4",
			args: args{
				buf: []byte{0x36, 0x55, 0x41, 0x46, 0x48, 0x72},
			},
			want: 0x34,
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
