// Provides stuttering functionality.

package uwu

import "strings"

// stutter takes a string and returns a stuttered version of it.
//
// The chance parameter is the chance of a word being stuttered.
// If a word is to be stuttered, it will be stuttered amount (0-100) amount of times.
func Stutter(text string, chance, amount int) string {
	if chance <= 0 || amount <= 0 {
		return text
	}

	var result strings.Builder

	words := strings.Split(text, " ")
	for i, word := range words {
		if rng.Intn(100) > chance {
			result.WriteString(word)
		} else {
			result.WriteString(stutterWord(word, amount))
		}

		// If this is the last word, we shouldn't add an extra space.
		if i != len(words)-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

// stutterWord takes a string and returns a stuttered version of it.
// The amount parameter is the amount (0-100) of times the word should be stuttered.
func stutterWord(word string, amount int) string {
	var result strings.Builder

	for i := 0; i < amount; i++ {
		result.WriteString(string(word[0]) + "- ")
	}

	result.WriteString(word)

	return result.String()
}
