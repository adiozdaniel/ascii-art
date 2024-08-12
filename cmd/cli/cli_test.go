package main

import (
	"testing"
	"time"
)

// TestRunCli tests the runCli function
func TestRunCli(t *testing.T) {
	done := make(chan bool)
	go func() {
		err := runCli()
		if err != nil {
			t.Errorf("runCli() error = %v, want nil", err)
		}
		done <- true
	}()

	time.Sleep(5 * time.Second)

	<-done
}
