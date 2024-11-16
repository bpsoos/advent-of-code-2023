package day2

import "io"

type GameVerifier struct{}

func (GameVerifier) Verify(games io.Reader) []int {
	return []int{}
}
