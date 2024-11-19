package day7

import "slices"

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
	return len(h.uniques()) == 4
}

func (h Hand) isTwoPair() bool {
	return len(h.uniques()) == 3
}

func (h Hand) isThreeOfAKind() bool {
	uniques := h.uniques()

	if len(uniques) != 3 {
		return false
	}

	for i := range uniques {
		if h.count(uniques[i]) == 3 {
			return true
		}
	}

	return false
}

func (h Hand) isFullHouse() bool {
	uniques := h.uniques()
	if len(uniques) != 2 {
		return false
	}

	count1 := h.count(uniques[0])
	count2 := h.count(uniques[1])

	return (count1 == 2 && count2 == 3) || (count1 == 3 && count2 == 2)
}

func (h Hand) isFourOfAKind() bool {
	for i := 0; i < 5; i++ {
		if h.count(h[i]) == 4 {
			return true
		}
	}

	return false
}

func (h Hand) isFiveOfAKind() bool {
	return h.count(h[0]) == 5
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
