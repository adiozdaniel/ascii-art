package utils

import (
	"testing"
)

// TestAlignment tests Alignment
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

	Inputs.Justify = "left"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alignment(tt.args.output, tt.args.width); len(got) != tt.want {
				t.Errorf("Alignment() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

// TestRemoveQuotes tests RemoveQuotes
func TestRemoveQuotes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "single quotes", args: args{input: "'hello'"}, want: "hello"},
		{name: "double quotes", args: args{input: "\"hello\""}, want: "hello"},
		{name: "quotes within", args: args{input: "hello \"world"}, want: "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveQuotes(tt.args.input); got != tt.want {
				t.Errorf("RemoveQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
