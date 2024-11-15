package tests

import (
	"testing"
)

// TestRemoveQuotes tests RemoveQuotes
func TestRemoveQuotes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "single quotes", args: args{input: "'hello'"}, want: "hello"},
		{name: "double quotes", args: args{input: "\"hello\""}, want: "hello"},
		{name: "quotes within", args: args{input: "hello \"world"}, want: "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := app.RemoveQuotes(tt.args.input); got != tt.want {
				t.Errorf("RemoveQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
