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

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand [5]Card

func (h Hand) Type() HandType {
	if h.isFiveOfAKind() {
		return FiveOfAKind
	}
	return HighCard
}

func (h Hand) isFiveOfAKind() bool {
	firstType := h[0]

	for i := 1; i < 5; i++ {
		if h[i] != firstType {
			return false
		}
	}

	return true
}

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

func (g Games) Swap(i, j int) {
	tmp := g[i]
	g[i] = g[j]
	g[j] = tmp
}

func (g Games) Len() int {
	return len(g)
}

func (g Games) Less(i, j int) bool {
	return false
}
