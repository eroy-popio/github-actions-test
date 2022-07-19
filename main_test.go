package main

import "testing"

func TestNewString(t *testing.T) {
	if NewString("Hello, World! ") != "Hello, World! 12345"{
		t.Error("Not the expected test output")
	}
}