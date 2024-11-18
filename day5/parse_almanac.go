package day5

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ParseAlmanac(rawAlmanac io.Reader) *Almanac {
	scanner := bufio.NewScanner(rawAlmanac)
	almanac := NewAlmanac()

	scanner.Scan()
	almanac.Seeds = getSeeds(scanner.Text())

	for scanner.Scan() {
		line := []rune(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[len(line)-1] == ':' {
			almanac.Maps = append(almanac.Maps, make(AlmanacMap, 0))
			continue
		}
		almanac.Maps[len(almanac.Maps)-1] = append(almanac.Maps[len(almanac.Maps)-1], getAlmanacMapEntry(line))
	}

	return almanac
}

func getSeeds(line string) []int {
	rawSeeds := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
	seeds := make([]int, len(rawSeeds))

	for i := 0; i < len(rawSeeds); i++ {
		seed, err := strconv.Atoi(rawSeeds[i])
		if err != nil {
			panic(err)
		}
		seeds[i] = seed
	}

	return seeds
}

func getAlmanacMapEntry(line []rune) AlmanacMapEntry {
	rawNumbers := strings.Split(strings.TrimPrefix(string(line), "seeds: "), " ")
	numbers := [3]int{}

	for i := 0; i < len(rawNumbers); i++ {
		number, err := strconv.Atoi(rawNumbers[i])
		if err != nil {
			panic(err)
		}
		numbers[i] = number
	}

	return numbers
}

type Almanac struct {
	Seeds             []int
	Maps              []AlmanacMap
	currentSeed       int
	currentRangedSeed int
}

func NewAlmanac() *Almanac {
	return &Almanac{
		Seeds: make([]int, 0),
		Maps:  make([]AlmanacMap, 0),
	}
}

func (a *Almanac) RangedSeed() (int, error) {
	currentRange := a.Seeds[a.currentSeed+1]
	if a.currentRangedSeed >= currentRange {
		a.currentRangedSeed = 0
		a.currentSeed += 2
	}

	if a.currentSeed >= len(a.Seeds) {
		return 0, errors.New("end of range")
	}

	seed := a.Seeds[a.currentSeed] + a.currentRangedSeed
	a.currentRangedSeed++

	return seed, nil
}

type AlmanacMap []AlmanacMapEntry

type AlmanacMapEntry [3]int

func (e AlmanacMapEntry) DstRangeStart() int {
	return e[0]
}

func (e AlmanacMapEntry) DstRangeEnd() int {
	return e[0] + e[2]
}

func (e AlmanacMapEntry) SrcRangeStart() int {
	return e[1]
}

func (e AlmanacMapEntry) SrcRangeEnd() int {
	return e[1] + e[2]
}

func (e AlmanacMapEntry) Offset() int {
	return e[0] - e[1]
}

func (e AlmanacMapEntry) Range() int {
	return e[2]
}
