// Package uwu provides functions to uwuify text.
//
// The most notable function is UwUify, which takes a string and
// returns a uwuified version of it.
package uwu

import "strings"

// A Config is a configuration for the UwUify function.
type Config struct {
	Lowercase      bool `query:"lower"`
	ReplaceLetters bool `query:"replace"`
	FacesChance    int  `query:"facesChance"`
	StutterChance  int  `query:"stutterChance"`
	StutterAmount  int  `query:"stutterAmount"`
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
