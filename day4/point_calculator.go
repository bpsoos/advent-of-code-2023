package day4

type PointCalculator struct{}

func (PointCalculator) Calculate(scratchcards Scratchcards) int {
	points := 0

	for i := 0; i < len(scratchcards); i++ {
		points += getPoints(&scratchcards[i])
	}

	return points
}

func (PointCalculator) CalculateWithCopies(scratchcards Scratchcards) int {
	originalCards := make(Scratchcards, len(scratchcards))
	copy(originalCards, scratchcards)

	stackLen := len(scratchcards)
	for i := 0; i < stackLen; i++ {
		count := getCopyCount(&scratchcards[i])
		stackLen += count
		for j := 0; j < count; j++ {
			scratchcards = append(scratchcards, originalCards[scratchcards[i].Id+j])
		}
	}

	return stackLen
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

func getCopyCount(scratchcard *Scratchcard) int {
	count := 0
	for i := 0; i < len(scratchcard.WinningNumbers); i++ {
		for j := 0; j < len(scratchcard.Numbers); j++ {
			if scratchcard.Numbers[j] != scratchcard.WinningNumbers[i] {
				continue
			}
			count++
		}
	}

	return count
}
