package gulu

// IsNumOrLetter
// checks the specified rune is number or letter.
func IsNumOrLetter(r rune) bool {
	return ('0' <= r && '9' >= r) || IsLetter(r)
}

// IsLetter
// checks the specified rune is letter.
func IsLetter(r rune) bool {
	return 'a' <= r && 'z' >= r || 'A' <= r && 'Z' >= r
}
