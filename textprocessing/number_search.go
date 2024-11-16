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
	endPos := 0
	for i := pos; i < len(chars); i++ {
		endPos = i
		if !unicode.IsDigit(chars[i]) {
			break
		}
	}
	value, err := strconv.Atoi(string(chars[pos:endPos]))
	if err != nil {
		panic(err)
	}

	return value, endPos - pos
}
