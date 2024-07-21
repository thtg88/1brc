package runes

import "unicode/utf8"

func RuneAt(s string, i uint) (rune, bool) {
	leftStr := s
	var j uint
	for {
		result, size := utf8.DecodeRuneInString(leftStr)
		if size == 0 {
			// we've reached the end of the string
			return 0, false
		}

		if j == i {
			// we've found the character!
			return result, true
		}

		leftStr = leftStr[size:]

		j++
	}
}
