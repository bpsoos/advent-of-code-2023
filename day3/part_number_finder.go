package day3

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

type PartNumberFinder struct{}
