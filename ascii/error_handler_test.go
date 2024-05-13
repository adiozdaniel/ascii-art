package main

import (
	"testing"

	utils "github.com/adiozdaniel/ascii-art/utilities"
)

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
