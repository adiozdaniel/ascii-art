package main

import (
	"reflect"
	"strings"
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
)

// TestOutput tests if the output by checking the length
func TestOutput(t *testing.T) {
	fileContent := FileContentTests()
	tests := []tests{
		{name: "ProcessInputWithHelloGoogle", args: args{color: "White", reff: "", fileContents: fileContent, input: []string{"Hello", "Google"}}, expected: len(asciiArt)},
		{name: "ProcessInputEmptyInput", args: args{color: "White", reff: "", fileContents: strings.Split("", "\n"), input: []string{}}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := len(ascii.Output(tt.args.color, tt.args.reff, tt.args.input, tt.args.fileContents))
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Output() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
