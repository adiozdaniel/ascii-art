package main

import (
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
)

func TestOutput(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ascii.Output(tt.args.input, tt.args.fileContents)
		})
	}
}
