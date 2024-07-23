package tests

import (
	"testing"

	"github.com/adiozdaniel/ascii-art/utils"
)

func TestAlignment(t *testing.T) {
	type args struct {
		output string
		width  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "long input", args: args{output: "abcdefghijklmnopqrstuvwxyz", width: 80}, want: 26},
		{name: "short input", args: args{output: "abcdefghijklm nnopqrstuvwxyz", width: 80}, want: 28},
		{name: "single character", args: args{output: "a", width: 80}, want: 1},
	}

	utils.Inputs.Justify = "left"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Alignment(tt.args.output, tt.args.width); len(got) != tt.want {
				t.Errorf("Alignment() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
