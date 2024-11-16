package textprocessing

import (
	"errors"
	"strconv"
	"unicode"
)

var digitWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func GetLetterDigit(chars []rune, pos int) (int, error) {
	if pos >= len(chars) || pos < 0 {
		return 0, errors.New("out of range")
	}
	for digitWord := range digitWords {
		compareEndPos := pos + len(digitWord)
		if compareEndPos > len(chars) {
			continue
		}
		compareStr := string(chars[pos:compareEndPos])
		if compareStr == digitWord {
			return digitWords[digitWord], nil
		}
	}

	return 0, errors.New("not found")
}

func GetLetterDigitBackward(chars []rune, pos int) (int, error) {
	if pos > len(chars) || pos < 2 {
		return 0, errors.New("out of range")
	}
	for digitWord := range digitWords {
		compareStartPos := pos - len(digitWord) + 1
		if compareStartPos < 0 {
			continue
		}
		compareStr := string(chars[compareStartPos : pos+1])
		if compareStr == digitWord {
			return digitWords[digitWord], nil
		}
	}

	return 0, errors.New("not found")
}

func MustGetLastDigit(chars []rune) int {
	for i := len(chars) - 1; i >= 0; i-- {
		if unicode.IsDigit(chars[i]) {
			digit, err := strconv.Atoi(string(chars[i]))
			if err != nil {
				panic(err)
			}
			return digit
		}
		digit, err := GetLetterDigitBackward(chars, i)
		if err == nil {
			return digit
		}
	}
	panic("not found")
}

func MustGetFirstDigit(chars []rune) int {
	for i := 0; i < len(chars); i++ {
		if unicode.IsDigit(chars[i]) {
			digit, err := strconv.Atoi(string(chars[i]))
			if err != nil {
				panic(err)
			}
			return digit
		}
		digit, err := GetLetterDigit(chars, i)
		if err == nil {
			return digit
		}
	}
	panic("not found")
}
