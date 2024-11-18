package day4

type PointCalculator struct{}

func (PointCalculator) Calculate(scratchcards Scratchcards) int {
	points := 0

	for i := 0; i < len(scratchcards); i++ {
		points += getPoints(&scratchcards[i])
	}

	return points
}
func getPoints(scratchcard *Scratchcard) int {
	points := 0
	for i := 0; i < len(scratchcard.WinningNumbers); i++ {
		for j := 0; j < len(scratchcard.Numbers); j++ {
			if scratchcard.Numbers[j] != scratchcard.WinningNumbers[i] {
				continue
			}
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
	}

	return points
}
