package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Horus")

	got := buf.String()
	want := "Hello, Horus"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
