package stringutil

import "strings"

// GetSpecialCharactersString returns a string containing all special characters
func GetSpecialCharactersString() string {
	return "!\"#$%&'()*,:;<=>?[\\]^`{|}-"
}

// ContainsSpecialCharacters returns true if the source string contains a special character
func ContainsSpecialCharacters(src string) bool {
	if strings.ContainsAny(src, GetSpecialCharactersString()) {
		return true
	}

	return false
}
