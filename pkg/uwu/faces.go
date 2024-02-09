// Provides Unicode face adding functionality.

package uwu

import "strings"

// Faces is a list of cute Unicode faces.
var Faces = []string{
	// This gun represents how much I hate this fucking project

	"(・`ω´・)", ";;w;;", "owo", "UwU", ">w<", "^w^", ">.<", "(* ^ ω ^)",
	"(⌒ω⌒)", "ヽ(*・ω・)ﾉ", "(o´∀`o)", "(o･ω･o)", "＼(＾▽＾)／", "(*^ω^)",
	"(◕‿◕✿)", "(◕ᴥ◕)", "ʕ•ᴥ•ʔ", "ʕ￫ᴥ￩ʔ", "(*^.^*)", "(｡♥‿♥｡)",
	"(◕ᴥ◕)", "ʕ•ᴥ•ʔ", "(✿ ♡‿♡)",
	"ʕ•ᴥ•ʔ", "(ᵔᴥᵔ)",
	"(• o •)", `">w<`,
	"ʕ￫ᴥ￩ʔ", "(*^ω^*)",
}

// AddFaces takes a string and returns a version of it with cute Unicode faces.
// chance (0-100) is the chance to add a face to a word.
func AddFaces(text string, chance int) string {
	var result strings.Builder

	words := strings.Split(text, " ")
	for i, word := range words {
		if rng.Intn(100) < chance {
			result.WriteString(word + " " + randomFace())
		} else {
			result.WriteString(word)
		}

		// If this is the last word, we shouldn't add an extra space.
		if i != len(words)-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

// randomFace returns a random face from Faces.
func randomFace() string {
	idx := rng.Intn(len(Faces))
	return Faces[idx]
}
