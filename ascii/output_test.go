package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
)

func TestOutput(t *testing.T) {
	fileContent := fileContentTests()
	tests := []tests{
		{name: "ProcessInputWithHelloGoogle", args: args{fileContents: fileContent, input: []string{"Hello", "Google"}}, expected: 624},
		{name: "ProcessInputEmptyInput", args: args{fileContents: strings.Split("", "\n"), input: []string{}}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := len(ascii.Output(tt.args.input, tt.args.fileContents))
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Output() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func fileContentTests() []string {
	fileName := "../data/standard.txt"
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fileContents := strings.Split(string(contents), "\n")
	return fileContents
}
