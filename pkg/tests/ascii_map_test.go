package tests

import (
	"reflect"
	"strings"
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/internals/ascii"
)

// TestAsciiMap tests fileContents output
func TestAsciiMap(t *testing.T) {
	tests := []tests{
		{name: "HelloGoogle", args: args{fileContents: strings.Split(asciiArt, "\n"), input: "Hello\nGoogle"}, expected: map[rune]int{32: 17}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ascii.AsciiMap(tt.args.fileContents)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("AsciiMap() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
