package day3

import (
	"advent2023/mathx"
	"advent2023/textprocessing"
	"bufio"
	"io"
	"unicode"
)

func ParseSchematic(schmaticReader io.Reader) *Schematic {
	scanner := bufio.NewScanner(schmaticReader)
	schematic := NewSchematic()
	y := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		x := 0
		for x < len(line) {
			if line[x] == '.' {
				x++
				continue
			}
			if unicode.IsDigit(line[x]) {
				value, length := textprocessing.MustGetNumber(line, x)
				schematic.PartNumbers = append(schematic.PartNumbers, PartNumber{
					Value: value,
					Position: PartNumberPosition{
						Start: mathx.Point{X: x, Y: y},
						End:   mathx.Point{X: x + length - 1, Y: y},
					},
				})
				x += length
				continue
			}
			schematic.AddSymbol(mathx.Point{X: x, Y: y})
			x++
		}
		y++
	}

	return schematic
}
