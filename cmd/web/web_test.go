package main

import "testing"

func TestRunWeb(t *testing.T) {
	err := runWeb()
	if err != nil {
		t.Errorf("runWeb() error = %v, want nil", err)
	}
}
