package main

import "testing"

func TestGetHelloString(t *testing.T) {
	expect := "Hello Germany!"
	got := getHelloString()
	if got != expect {
		t.Errorf("\nExpect: %s \nGot: %s", expect, got)
	}
}
