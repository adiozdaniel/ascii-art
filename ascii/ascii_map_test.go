package main

import (
	"reflect"
	"testing"
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
