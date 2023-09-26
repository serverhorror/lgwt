package main

import "testing"

func TestHelloWho(t *testing.T) {
	got := HelloWho("World")
	want := "Hello World!"
	if got != want {
		t.Errorf("HelloWho() = %q, want %q", got, want)
	}
}

func TestHelloWorld(t *testing.T) {
	want := "Hello World!"
	got := HelloWorld()
	if got != want {
		t.Errorf("HelloWorld() = %q, want %q", got, want)
	}
}
