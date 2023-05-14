// Tests for the AddFaces function.

package uwu

import "testing"

// TestNoFaces tests the AddFaces function with a 0% chance of adding faces.
func TestNoFaces(t *testing.T) {
	t.Parallel()

	text := "Hello, world!"
	expected := "Hello, world!"
	actual := AddFaces(text, 0)

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

// TestFullFaces tests the AddFaces function with a 100% chance of adding faces.
func TestFullFaces(t *testing.T) {
	t.Parallel()

	fakeRNG()

	text := "Hello, world!"
	expected := "Hello, (✿ ♡‿♡) world! (o´∀`o)"
	actual := AddFaces(text, 100)

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

// TestHalfFaces tests the AddFaces function with a 50% chance of adding faces.
func TestHalfFaces(t *testing.T) {
	t.Parallel()

	fakeRNG()

	text := "Hello, world!"
	expected := "Hello, world! (◕ᴥ◕)"
	actual := AddFaces(text, 50)

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
