package main

import (
	"strings"
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
)

func TestOutput(t *testing.T) {
	tests := []struct {
		name string
		args Args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ascii.Output(strings.Split(tt.args.input, "\\n"), tt.args.fileContents)
		})
	}
}
