package main

import (
	"testing"

	utils "github.com/adiozdaniel/ascii-art/utilities"
)

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
