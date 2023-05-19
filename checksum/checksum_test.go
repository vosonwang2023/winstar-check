package checksum

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
				buf: []byte{0x01, 0x02, 0x03, 0x04},
			},
			want: 0x0A,
		},
		{
			name: "TestCalc2",
			args: args{
				buf: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			},
			want: 0x0F,
		},
		{
			name: "TestCalc3",
			args: args{
				buf: []byte{0x33, 0x35, 0x37, 0x45},
			},
			want: 0xE4,
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
