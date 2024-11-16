package day2

type MinSetFinder struct{}

type MinSets []CubeSet

func (MinSetFinder) Find(games []Game) MinSets {
	minSets := make(MinSets, 0)
	for _, game := range games {
		minSet := CubeSet{}
		for _, set := range game.RevealedSets {
			if set.Red > minSet.Red {
				minSet.Red = set.Red
			}
			if set.Green > minSet.Green {
				minSet.Green = set.Green
			}
			if set.Blue > minSet.Blue {
				minSet.Blue = set.Blue
			}
		}
		minSets = append(minSets, minSet)
	}

	return minSets
}
