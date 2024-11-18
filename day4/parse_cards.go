package day4

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Scratchcards []Scratchcard

type Scratchcard struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
}

func ParseCards(rawCards io.Reader) Scratchcards {
	scanner := bufio.NewScanner(rawCards)
	scratchcards := make(Scratchcards, 0)

	for scanner.Scan() {
		line := scanner.Text()
		scratchcards = append(scratchcards, Scratchcard{
			Id:             getCardId(line),
			WinningNumbers: getWinningNumbers(line),
			Numbers:        getNumbers(line),
		})
	}

	return scratchcards
}

func getCardId(line string) int {
	rawId := strings.TrimSpace(strings.TrimPrefix(strings.Split(line, ":")[0], "Card"))
	id, err := strconv.Atoi(rawId)
	if err != nil {
		panic(err)
	}
	return id
}
func getWinningNumbers(line string) []int {
	rawNumbersUntrimmed := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[0]), " ")
	rawNumbers := make([]string, 0)
	for _, rawNumberUntrimmed := range rawNumbersUntrimmed {
		rawNumber := strings.TrimSpace(rawNumberUntrimmed)
		if rawNumber != "" {
			rawNumbers = append(rawNumbers, rawNumber)
		}
	}
	winningNumbers := make([]int, len(rawNumbers))

	for i := 0; i < len(rawNumbers); i++ {
		number, err := strconv.Atoi(rawNumbers[i])
		if err != nil {
			panic(err)
		}
		winningNumbers[i] = number
	}

	return winningNumbers
}
func getNumbers(line string) []int {
	rawNumbersUntrimmed := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[1]), " ")
	rawNumbers := make([]string, 0)
	for _, rawNumberUntrimmed := range rawNumbersUntrimmed {
		rawNumber := strings.TrimSpace(rawNumberUntrimmed)
		if rawNumber != "" {
			rawNumbers = append(rawNumbers, rawNumber)
		}
	}
	numbers := make([]int, len(rawNumbers))

	for i := 0; i < len(rawNumbers); i++ {
		number, err := strconv.Atoi(rawNumbers[i])
		if err != nil {
			panic(err)
		}
		numbers[i] = number
	}

	return numbers
}
