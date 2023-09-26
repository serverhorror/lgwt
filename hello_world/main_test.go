package main

import (
	"testing"
)

func TestI18NHelloWho(t *testing.T) {
	testCases := []struct {
		who     string
		lang    string
		want    string
		wantErr bool
	}{
		{"", "en", "Hello World!", false},
		{"", "es", "Hola World!", false},
		{"", "fr", "Bonjour World!", false},
		{"", "de", "Hello World!", true},
		{"Alice", "en", "Hello Alice!", false},
		{"Alice", "es", "Hola Alice!", false},
		{"Alice", "fr", "Bonjour Alice!", false},
		{"Alice", "de", "Hello Alice!", true},
	}
	for _, tc := range testCases {
		got, err := HelloI18nWho(tc.who, tc.lang)
		assertCorrectMessageWithErr(t, got, tc.want, err, tc.wantErr)
	}
}

func TestHelloWho(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWho("Chris")
		want := "Hello Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloWho("")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessageWithErr(t testing.TB, got, want string, err error, wantErr bool) {
	t.Helper()
	if err != nil && !wantErr {
		t.Errorf("got %q wantErr %v", err, wantErr)
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
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
