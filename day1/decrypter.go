package day1

import (
	"advent2023/math"
	"bufio"
	"io"
	"strconv"
	"unicode"
)

type Decrypter struct{}

func (d Decrypter) Decrypt(calibrationDocument io.Reader) int {
	scanner := bufio.NewScanner(calibrationDocument)
	results := make([]int, 0)
	for scanner.Scan() {
		var (
			firstDigit, lastDigit int
			err                   error
		)
		chars := []rune(scanner.Text())
		for _, char := range chars {
			if unicode.IsDigit(char) {
				firstDigit, err = strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}
				break
			}
		}
		for i := len(chars) - 1; i >= 0; i-- {
			if unicode.IsDigit(chars[i]) {
				lastDigit, err = strconv.Atoi(string(chars[i]))
				if err != nil {
					panic(err)
				}
				break
			}
		}
		results = append(results, firstDigit*10+lastDigit)
	}

	return math.Sum(results)
}
