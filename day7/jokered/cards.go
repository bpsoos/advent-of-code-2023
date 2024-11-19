package jokered

import "slices"

type Card int

const (
	Joker Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Queen
	King
	Ace
)

func (c Card) String() string {
	switch c {
	case Joker:
		return "Joker"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		panic("unknown card")
	}
}

type HandType int

func (ht HandType) String() string {
	switch ht {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	default:
		panic("unknown hand type")
	}
}

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
	if h.isFourOfAKind() {
		return FourOfAKind
	}
	if h.isFullHouse() {
		return FullHouse
	}
	if h.isThreeOfAKind() {
		return ThreeOfAKind
	}
	if h.isTwoPair() {
		return TwoPair
	}
	if h.isOnePair() {
		return OnePair
	}

	return HighCard
}

func (h Hand) isOnePair() bool {
	regulars := h.regulars()
	if len(regulars) < 4 {
		return false
	}
	uniques := h.uniques()
	if len(regulars) == 4 {
		return len(uniques) == 5
	}

	return len(uniques) == 4
}

func (h Hand) isTwoPair() bool {
	regulars := h.regulars()
	if len(regulars) != 5 {
		return false
	}

	return len(h.uniques()) == 3
}

func (h Hand) isThreeOfAKind() bool {
	regulars := h.regulars()
	if len(regulars) < 3 {
		return false
	}

	uniques := h.uniques()
	if len(regulars) != 5 {
		return len(uniques) == 4
	}

	for i := range uniques {
		if h.count(uniques[i]) == 3 {
			return true
		}
	}

	return false
}

func (h Hand) isFullHouse() bool {
	regulars := h.regulars()
	if len(regulars) < 4 {
		return false
	}
	if len(regulars) == 4 {
		return len(h.uniques()) == 3
	}
	uniques := h.uniques()
	if len(uniques) != 2 {
		return false
	}

	count1 := h.count(uniques[0])
	count2 := h.count(uniques[1])

	return (count1 == 2 && count2 == 3) || (count1 == 3 && count2 == 2)
}

func (h Hand) isFourOfAKind() bool {
	regulars := h.regulars()
	uniques := h.uniques()
	jokerCount := 5 - len(regulars)

	if jokerCount > 3 {
		return false
	}
	if jokerCount == 3 {
		return true
	}
	if jokerCount == 2 {
		return len(uniques) == 3
	}

	if jokerCount == 1 {
		if len(uniques) != 3 {
			return false
		}
		return h.count(uniques[0]) == 3 || h.count(uniques[1]) == 3 || h.count(uniques[2]) == 3
	}

	if len(uniques) != 2 {
		return false
	}
	return h.count(uniques[0]) == 4 || h.count(uniques[1]) == 4
}

func (h Hand) isFiveOfAKind() bool {
	regulars := h.regulars()
	if len(regulars) == 0 {
		return true
	}
	return h.count(regulars[0]) == len(regulars)
}

func (h Hand) regulars() []Card {
	regulars := make([]Card, 0)

	for i := range 5 {
		if h[i] != Joker {
			regulars = append(regulars, h[i])
		}
	}

	return regulars
}

func (h Hand) uniques() []Card {
	uniques := make([]Card, 0)

	for i := range 5 {
		if !slices.Contains(uniques, h[i]) {
			uniques = append(uniques, h[i])
		}
	}

	return uniques
}

func (h Hand) count(card Card) int {
	count := 0

	for i := range 5 {
		if h[i] == card {
			count++
		}
	}

	return count
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
	iType := g[i].Hand.Type()
	jType := g[j].Hand.Type()
	if iType != jType {
		return iType < jType
	}
	for k := range 5 {
		iCard := g[i].Hand[k]
		jCard := g[j].Hand[k]
		if iCard != jCard {
			return iCard < jCard
		}
	}
	panic("could not compare hands")
}
