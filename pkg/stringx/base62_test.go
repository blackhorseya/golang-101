package stringx

import (
	"testing"
)

func TestEncodeBase62(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{n: 0},
			want: "0",
		},
		{
			name: "1",
			args: args{n: 1},
			want: "1",
		},
		{
			name: "62",
			args: args{n: 62},
			want: "10",
		},
		{
			name: "3844",
			args: args{n: 3844},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase62(tt.args.n); got != tt.want {
				t.Errorf("EncodeBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase62(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "DecodeBase62 for 0",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "DecodeBase62 for 1",
			args: args{s: "1"},
			want: 1,
		},
		{
			name: "DecodeBase62 for 10",
			args: args{s: "10"},
			want: 62,
		},
		{
			name: "DecodeBase62 for 100",
			args: args{s: "100"},
			want: 3844,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeBase62(tt.args.s); got != tt.want {
				t.Errorf("DecodeBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}
