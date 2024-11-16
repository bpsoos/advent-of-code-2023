package textprocessing

import (
	"errors"
	"strconv"
	"unicode"
)

func MustGetNumber(chars []rune, pos int) (value int, length int) {
	if pos >= len(chars) {
		panic(errors.New("out of range"))
	}
	endPos := pos + 1
	for endPos < len(chars) {
		isDigit := unicode.IsDigit(chars[endPos])
		if !isDigit {
			break
		}
		endPos++
	}
	value, err := strconv.Atoi(string(chars[pos:endPos]))
	if err != nil {
		panic(err)
	}

	return value, endPos - pos
}
