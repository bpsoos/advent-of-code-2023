package day7

import "sort"

type TotalWinningsCalculator struct{}

func (twc TotalWinningsCalculator) Calculate(games Games) int {
	sort.Sort(games)
	return twc.sumBids(games)
}

func (TotalWinningsCalculator) sumBids(games Games) int {
	return 0
}
