package day1

import "os"

type Solver struct{}

func (solver Solver) Solve() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
    println("day 1 solution:")
	println(Decrypter{}.Decrypt(f))
}
