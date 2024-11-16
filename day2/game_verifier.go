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
			if !g.isPossible(set) {
				continue loop
			}
		}
		possibleGames = append(possibleGames, game.Id)
	}
	return possibleGames
}

func (g GameVerifier) isPossible(set CubeSet) bool {
	return set.Red < g.BagContent.Red && set.Blue < g.BagContent.Blue && set.Green < g.BagContent.Green
}
