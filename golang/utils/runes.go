package utils

func IsLowerCase(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpperCase(c rune) bool {
	return c >= 'A' && c <= 'Z'
}
