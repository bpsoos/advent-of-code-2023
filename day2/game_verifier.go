package day2

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id           int
	RevealedSets []CubeSet
}

type GameVerifier struct {
	BagContent CubeSet
}

func (g GameVerifier) Verify(games []Game) []int {
	possibleGames := make([]int, 0)
loop:
	for _, game := range games {
		for _, set := range game.RevealedSets {
			if set.Red > g.BagContent.Red {
				continue loop
			}
			if set.Blue > g.BagContent.Blue {
				continue loop
			}
			if set.Green > g.BagContent.Green {
				continue loop
			}
		}
		possibleGames = append(possibleGames, game.Id)
	}
	return possibleGames
}
