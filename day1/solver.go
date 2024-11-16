package day1

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	contents, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 1 solution 1:")
	fmt.Println(Decrypter{}.Decrypt(bytes.NewReader(contents)))
	fmt.Println("day 1 solution 2:")
	fmt.Println(Decrypter{}.DecryptWithLetters(bytes.NewReader(contents)))
}
