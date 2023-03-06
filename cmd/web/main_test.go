package main

import "testing"

func TestRun(t *testing.T) {
	if err := run(); err != nil {
		t.Errorf("Error in run function: %v", err)
	}

}
