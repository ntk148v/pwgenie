package main

import "testing"

func Test_say(t *testing.T) {
	want := "Expected"
	if got := say("Expected"); got != want {
		t.Errorf("say() = %v, want %v", got, want)
	}
}
