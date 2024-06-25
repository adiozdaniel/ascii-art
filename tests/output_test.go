package tests

import (
	"github.com/adiozdaniel/ascii-art/utils"
	"reflect"
	"strings"
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
)

// TestOutput tests if the output is meeting requirements by checking the length of its output
func TestOutput(t *testing.T) {
	fileContent := FileContentTests()
	tests := []tests{
		{name: "ProcessInputWithHelloGoogle", args: args{fileContents: fileContent, input: "Hello Google"}, expected: 8},
		{name: "ProcessInputEmptyInput", args: args{fileContents: strings.Split("", "\n"), input: ""}, expected: `🧐 Oops! We can't find your "something"

		EX: go run . [OPTION] [STRING] "something"
		`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.Inputs.Input = tt.args.input
			actual := len(ascii.Output(tt.args.fileContents))
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Output() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
