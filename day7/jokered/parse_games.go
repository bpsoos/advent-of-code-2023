package jokered

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseGames(rawGames io.Reader) Games {
	scanner := bufio.NewScanner(rawGames)
	games := make(Games, 0)

	for scanner.Scan() {
		games = append(games, *parseGame(scanner.Text()))
	}

	return games
}

func parseGame(line string) *Game {
	splitline := strings.Split(strings.TrimSpace(line), " ")
	cards := []rune(splitline[0])
	rawBid := splitline[1]
	game := NewGame()

	if len(cards) != 5 {
		panic("card count must be 5")
	}

	for i := range 5 {
		game.Hand[i] = parseCard(cards[i])
	}

	bid, err := strconv.Atoi(rawBid)
	if err != nil {
		panic(err)
	}
	game.Bid = bid

	return game
}

func parseCard(char rune) Card {
	switch char {
	case '2':
		return Two
	case '3':
		return Three
	case '4':
		return Four
	case '5':
		return Five
	case '6':
		return Six
	case '7':
		return Seven
	case '8':
		return Eight
	case '9':
		return Nine
	case 'T':
		return Ten
	case 'J':
		return Joker
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		return Ace
	default:
		panic("unknown card")
	}
}
