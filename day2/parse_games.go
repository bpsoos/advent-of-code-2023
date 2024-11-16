package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseGames(rawGames io.Reader) []Game {
	games := make([]Game, 0)
	scanner := bufio.NewScanner(rawGames)
	for scanner.Scan() {
		rawGame := scanner.Text()
		rawGame = strings.TrimSpace(rawGame)
		if rawGame == "" {
			continue
		}
		gameId := getGameId(rawGame)
		game := Game{Id: gameId}
		game.RevealedSets = getRevealedCubeSets(rawGame)
		games = append(games, game)
	}
	return games
}

func getGameId(rawGame string) int {
	gameInfo := strings.Split(rawGame, ":")[0]
	rawGameId := strings.Split(gameInfo, " ")[1]
	gameId, err := strconv.Atoi(rawGameId)
	if err != nil {
		panic(err)
	}
	return gameId
}

func getRevealedCubeSets(rawGame string) []CubeSet {
	sets := make([]CubeSet, 0)

	for _, rawSet := range getRawSets(rawGame) {
		set := new(CubeSet)

		for _, cubeByColor := range getCubesByColor(rawSet) {
			value := getValueForCube(cubeByColor)
			switch color := getColorForCube(cubeByColor); color {
			case "r":
				set.Red = value
			case "g":
				set.Green = value
			case "b":
				set.Blue = value
			default:
				panic("unknown color")
			}
		}
		sets = append(sets, *set)
	}
	return sets
}

func getValueForCube(cubeByColor string) int {
	rawCubeValue := strings.Split(strings.TrimSpace(cubeByColor), " ")[0]
	value, err := strconv.Atoi(rawCubeValue)
	if err != nil {
		panic(err)
	}
	return value
}

func getRawSets(rawGame string) []string {
	return strings.Split(strings.Split(rawGame, ":")[1], ";")
}

func getCubesByColor(rawSet string) []string {
	return strings.Split(strings.TrimSpace(rawSet), ",")
}

func getColorForCube(cubeByColor string) string {
	return string(strings.Split(strings.TrimSpace(cubeByColor), " ")[1][0])
}
