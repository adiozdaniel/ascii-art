package main

import (
	"reflect"
	"strings"
	"testing"

	ascii "github.com/adiozdaniel/ascii-art/art_work"
	utils "github.com/adiozdaniel/ascii-art/utilities"
)

func TestAsciiMap(t *testing.T) {
	type args struct {
		fileContents []string
	}
	tests := []struct {
		name     string
		args     args
		expected map[rune]int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.args.fileContents); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("AsciiMap() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestGetFile(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetFile(tt.args.args); got != tt.expected {
				t.Errorf("GetFile() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	type args struct {
		input        string
		fileContents []string
		asciiMap     map[rune]int
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ascii.Output(strings.Split(tt.args.input, "\\n"), tt.args.fileContents)
		})
	}
}

func TestErrorHandler(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.ErrorHandler(tt.args.err)
		})
	}
}
