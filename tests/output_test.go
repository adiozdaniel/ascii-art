package tests

import (
	"reflect"
	"strings"
	"testing"

	"github.com/adiozdaniel/ascii-art/utils"

	ascii "github.com/adiozdaniel/ascii-art/ascii"
)

// TestOutput tests if the output is meeting requirements by checking the length of its output
func TestOutput(t *testing.T) {
	fileContent := FileContentTests()
	tests := []tests{
		{name: "ProcessInputWithHelloGoogle", args: args{fileContents: fileContent, input: "Hello\nGoogle"}, expected: 18},
		{name: "ProcessOneChar", args: args{fileContents: strings.Split("", "\n"), input: "1"}, expected: 9},
		{name: "ProcessInputEmptyInput", args: args{fileContents: strings.Split("", "\n"), input: ""}, expected: 0},
	}

	utils.Inputs.BannerPath = utils.BannerFiles["standard"]

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := len(ascii.Output(tt.args.input))
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Output() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
