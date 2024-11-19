package day7

import (
	"advent2023/day7/jokered"
	"fmt"
	"sort"
)

type TotalWinningsCalculator struct{}

func (twc TotalWinningsCalculator) Calculate(games Games) int {
	sort.Sort(games)
	return twc.sumBids(games)
}

func (twc TotalWinningsCalculator) CalculateJokered(games jokered.Games) int {
	sort.Sort(games)
	return twc.sumJokeredBids(games)
}

func (TotalWinningsCalculator) sumJokeredBids(games jokered.Games) int {
	sum := 0
	for i := 0; i < len(games); i++ {
		fmt.Printf("hand: %v\n", games[i].Hand)
		fmt.Printf("type: %v\n", games[i].Hand.Type())
		sum += games[i].Bid * (i + 1)
	}
	return sum
}

func (TotalWinningsCalculator) sumBids(games Games) int {
	sum := 0
	for i := 0; i < len(games); i++ {
		sum += games[i].Bid * (i + 1)
	}
	return sum
}
