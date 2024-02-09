// Replaces letters to make text sound "cute".

package uwu

import "strings"

// ReplaceLetters replaces letters to make text sound "cute".
// It replaces all instances of "r" and "l" with "w".
func ReplaceLetters(text string) string {
	text = strings.ReplaceAll(text, "r", "w")
	text = strings.ReplaceAll(text, "l", "w")
	text = strings.ReplaceAll(text, "R", "W")
	text = strings.ReplaceAll(text, "L", "W")

	return text
}
