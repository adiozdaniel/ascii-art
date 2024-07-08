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
		{name: "ProcessInputWithHelloGoogle", args: args{fileContents: fileContent, input: "Hello\nGoogle"}, expected: 626},
		{name: "ProcessOneChar", args: args{fileContents: strings.Split("", "\n"), input: "1"}, expected: 635},
		{name: "ProcessInputEmptyInput", args: args{fileContents: strings.Split("", "\n"), input: ""}, expected: 0},
	}

	utils.Inputs.BannerPath = utils.BannerFiles["standard"]

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
