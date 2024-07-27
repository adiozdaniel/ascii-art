package ascii

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/adiozdaniel/ascii-art/utils"
)

// Structs used to store the data in test files
type args struct {
	input        string
	fileContents []string
}

type tests struct {
	name     string
	args     args
	expected interface{}
}

// asciiArt is used as a test case
const asciiArt = `      _    _          _   _  ` + "\n" +
	`| |  | |        | | | |         ` + "\n" +
	`| |__| |   ___  | | | |   ___   ` + "\n" +
	`|  __  |  / _ \ | | | |  / _ \  ` + "\n" +
	`| |  | | |  __/ | | | | | (_) | ` + "\n" +
	`|_|  |_|  \___| |_| |_|  \___/  ` + "\n" +
	`                                  ` + "\n" +
	`                                  ` + "\n" +
	`  _____                           _         ` + "\n" +
	` / ____|                         | |        ` + "\n" +
	`| |  __    ___     ___     __ _  | |   ___  ` + "\n" +
	`| | |_ |  / _ \   / _ \   / _` + "`" + ` | | |  / _ \ ` + "\n" +
	`| |__| | | (_) | | (_) | | (_| | | | |  __/ ` + "\n" +
	` \_____|  \___/   \___/   \__, | |_|  \___| ` + "\n" +
	`                           __/ |            ` + "\n" +
	`                          |___/            ` + "\n"

// TestOutput tests if the output is meeting requirements by checking the length of its output
func TestOutput(t *testing.T) {
	fileContent := FileContentTests()
	tests := []tests{
		{name: "ProcessInputWithHelloGoogle", args: args{fileContents: fileContent, input: "Hello\nGoogle"}, expected: 9},
		{name: "ProcessOneChar", args: args{fileContents: strings.Split("", "\n"), input: "1"}, expected: 9},
		{name: "ProcessInputEmptyInput", args: args{fileContents: strings.Split("", "\n"), input: ""}, expected: 0},
	}

	utils.Inputs.BannerPath = utils.BannerFiles["standard"]

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := len(Output(tt.args.input))
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Output() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

// FileContentTests checks the content of a bannerfile
func FileContentTests() []string {
	fileName := "../banners/standard.txt"
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fileContents := strings.Split(string(contents), "\n")
	return fileContents
}

// TestAsciiMap tests fileContents output
func TestAsciiMap(t *testing.T) {
	tests := []tests{
		{name: "HelloGoogle", args: args{fileContents: strings.Split(asciiArt, "\n"), input: "Hello\nGoogle"}, expected: map[rune]int{32: 17}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := AsciiMap(tt.args.fileContents)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("AsciiMap() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
