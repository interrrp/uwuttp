// Package uwu provides functions to uwuify text.
//
// The most notable function is UwUify, which takes a string and
// returns a uwuified version of it.
package uwu

import "strings"

// A Config is a configuration for the UwUify function.
type Config struct {
	// Lowercase determines whether or not the text should be lowercased.
	Lowercase bool `query:"lower"`

	// ReplaceLetters determines whether or not letters should be replaced.
	ReplaceLetters bool `query:"replace"`

	// Faces determines whether or not cute Unicode faces should be added.
	FacesChance int `query:"facesChance"`

	// StutterChance is the chance of a word being stuttered.
	StutterChance int `query:"stutterChance"`

	// StutterAmount is the amount of times a word should be stuttered.
	StutterAmount int `query:"stutterAmount"`
}

// NewConfig returns a new Config with default values.
func NewConfig() Config {
	return Config{
		Lowercase:      true,
		ReplaceLetters: true,
		FacesChance:    10,
		StutterChance:  30,
		StutterAmount:  1,
	}
}

// UwUify takes a string and returns a uwuified version of it.
func UwUify(text string, cfg Config) string {
	if cfg.Lowercase {
		text = strings.ToLower(text)
	}

	if cfg.ReplaceLetters {
		text = ReplaceLetters(text)
	}

	text = Stutter(text, cfg.StutterChance, cfg.StutterAmount)
	text = AddFaces(text, cfg.FacesChance)

	return text
}
