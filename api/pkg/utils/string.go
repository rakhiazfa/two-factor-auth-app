package utils

import "unicode"

func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
