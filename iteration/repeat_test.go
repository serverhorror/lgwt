package iteration

import (
	"strings"
	"testing"
)

func Repeat(character string) string {
	var sb strings.Builder

	for i := 0; i < 5; i++ {
		sb.WriteString(character)
	}
	return sb.String()
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}
