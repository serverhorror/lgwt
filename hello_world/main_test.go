package main

import "testing"

func TestHelloWho(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWho("Chris")
		want := "Hello Chris!"
		assertCorretMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloWho("")
		want := "Hello World!"
		assertCorretMessage(t, got, want)
	})
}

func assertCorretMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWorld(t *testing.T) {
	want := "Hello World!"
	got := HelloWorld()
	if got != want {
		t.Errorf("HelloWorld() = %q, want %q", got, want)
	}
}
