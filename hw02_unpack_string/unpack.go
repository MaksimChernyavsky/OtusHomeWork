package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	sRunes := []rune(s)
	var sb strings.Builder
	for i := 0; i < len(sRunes); {
		switch {
		case unicode.IsDigit(sRunes[i]):
			return "", ErrInvalidString
		case i == len(sRunes)-1 || !unicode.IsDigit(sRunes[i+1]):
			sb.WriteRune(sRunes[i])
		case unicode.IsDigit(sRunes[i+1]):
			c, _ := strconv.Atoi(string(sRunes[i+1]))
			t := string(sRunes[i])
			sb.WriteString(strings.Repeat(t, c))
			i++
		}
		i++
	}
	return sb.String(), nil
}
