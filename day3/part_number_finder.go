package day3

import (
	"reflect"
)

type PartNumberFinder struct{}

func (PartNumberFinder) Find(schematic *Schematic) []int {
	partNumbers := make([]int, 0, len(schematic.PartNumbers))
	for _, partNumber := range schematic.PartNumbers {
		for _, adjPosition := range partNumber.AdjPositions() {
			if schematic.IsSymbol(adjPosition) {
				partNumbers = append(partNumbers, partNumber.Value)
				break
			}
		}

	}
	return partNumbers
}

func (PartNumberFinder) FindGearRatios(schematic *Schematic) []int {
	gearRatios := make([]int, 0, len(schematic.Gears))

	for _, gear := range schematic.Gears {
		gearRatios = append(gearRatios, getGearRatio(schematic, &gear))
	}

	return gearRatios
}

func getGearRatio(schematic *Schematic, gear *PartPosition) int {
	var (
		firstPart  *PartNumber
		secondPart *PartNumber
	)
	for _, adjPosition := range gear.AdjPositions() {
		part, err := schematic.FindPart(adjPosition)
		if err != nil {
			continue
		}
		if firstPart == nil || reflect.DeepEqual(part, firstPart) {
			firstPart = part
			continue
		}
		if secondPart == nil || reflect.DeepEqual(part, secondPart) {
			secondPart = part
			continue
		}
		return 0
	}

	if secondPart == nil || firstPart == nil {
		return 0
	}

	return firstPart.Value * secondPart.Value
}
