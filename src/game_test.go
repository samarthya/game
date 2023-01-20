// testing package game methods
package game

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world!"

	if got := Hello(); got != want {
		t.Errorf(" Hello() = %q, want %q", got, want)
	}
}

func TestQuote(t *testing.T) {
	want := "Hello, world."

	if got := Quote(); got != want {
		t.Errorf(" Quote() = %q, want %q", got, want)
	}
}
