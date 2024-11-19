package day7

type Card int

const (
	Two Card = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Hand [5]Card

type Game struct {
	Hand Hand
	Bid  int
}

func NewGame() *Game {
	return &Game{
		Hand: Hand{},
		Bid:  0,
	}
}

type Games []Game

func (g Games) Swap(i, j int)      {}
func (g Games) Len() int           { return 0 }
func (g Games) Less(i, j int) bool { return false }
