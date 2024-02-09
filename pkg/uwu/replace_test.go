// Tests for the ReplaceLetters function.

package uwu

import "testing"

// TestReplaceLettersEmpty tests the ReplaceLetters function with an empty string.
func TestReplaceLettersEmpty(t *testing.T) {
	if ReplaceLetters("") != "" {
		t.Error("ReplaceLetters doesn't work with an empty string")
	}
}

// TestReplaceLettersSingle tests the ReplaceLetters function with a single letter.
func TestReplaceLettersSingle(t *testing.T) {
	if ReplaceLetters("r") != "w" {
		t.Error("ReplaceLetters doesn't work with a single letter")
	}
}

// TestReplaceLettersMultiple tests the ReplaceLetters function with multiple letters.
func TestReplaceLettersMultiple(t *testing.T) {
	if ReplaceLetters("hello world") != "hewwo wowwd" {
		t.Error("ReplaceLetters doesn't work with multiple letters")
	}
}
