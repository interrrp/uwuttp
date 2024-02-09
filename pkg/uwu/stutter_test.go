// Tests for the stuttering functionality.

package uwu

import "testing"

// TestNoStutter tests the Stutter function with a 0% chance of stuttering.
func TestNoStutter(t *testing.T) {
	t.Parallel()

	text := "Hello, world!"
	expected := "Hello, world!"
	actual := Stutter(text, 0, 0)

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

// TestFullStutter tests the Stutter function with a 100% chance of stuttering.
func TestFullStutter(t *testing.T) {
	t.Parallel()

	text := "Hello, world!"
	expected := "H- Hello, w- world!"
	actual := Stutter(text, 100, 1)

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

// TestHalfStutter tests the Stutter function with a 50% chance of stuttering.
func TestHalfStutter(t *testing.T) {
	t.Parallel()

	fakeRNG()

	text := "Hello, world!"
	expected := "Hello, w- world!"
	actual := Stutter(text, 50, 1)

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
