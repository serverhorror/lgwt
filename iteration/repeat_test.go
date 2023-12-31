package iteration

import (
	"fmt"
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

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
