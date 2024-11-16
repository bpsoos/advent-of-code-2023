package day3

import (
	"advent2023/mathx"
	"strconv"
)

func NewSchematic() *Schematic {
	return &Schematic{
		symbolPositions: make(map[string]bool),
		PartNumbers:     make([]PartNumber, 0),
	}
}

type (
	Schematic struct {
		symbolPositions map[string]bool
		PartNumbers     []PartNumber
	}

	PartNumber struct {
		Value    int
		Position PartNumberPosition
	}

	PartNumberPosition mathx.LineSegment
)

func (pn PartNumber) AdjPositions() []mathx.Point {
	positionCount := 6 + pn.Position.GetLength()*2
	adjPositions := make([]mathx.Point, 0, positionCount)

	adjPositions = append(adjPositions, mathx.Point{X: pn.Position.Start.X - 1, Y: pn.Position.Start.Y - 1})
	adjPositions = append(adjPositions, mathx.Point{X: pn.Position.Start.X - 1, Y: pn.Position.Start.Y})
	adjPositions = append(adjPositions, mathx.Point{X: pn.Position.Start.X - 1, Y: pn.Position.Start.Y + 1})

	adjPositions = append(adjPositions, mathx.Point{X: pn.Position.End.X + 1, Y: pn.Position.End.Y - 1})
	adjPositions = append(adjPositions, mathx.Point{X: pn.Position.End.X + 1, Y: pn.Position.End.Y})
	adjPositions = append(adjPositions, mathx.Point{X: pn.Position.End.X + 1, Y: pn.Position.End.Y + 1})

	for i := 0; i < pn.Position.GetLength(); i++ {
		adjPositions = append(adjPositions, mathx.Point{X: pn.Position.Start.X + i, Y: pn.Position.Start.Y - 1})
		adjPositions = append(adjPositions, mathx.Point{X: pn.Position.Start.X + i, Y: pn.Position.Start.Y + 1})
	}

	return adjPositions
}
func (pnp PartNumberPosition) GetLength() int {
	return pnp.End.X - pnp.Start.X + 1
}

func (s *Schematic) AddSymbol(position mathx.Point) {
	positionKey := strconv.Itoa(position.X) + "," + strconv.Itoa(position.Y)
	s.symbolPositions[positionKey] = true
}

func (s *Schematic) IsSymbol(position mathx.Point) bool {
	positionKey := strconv.Itoa(position.X) + "," + strconv.Itoa(position.Y)
	_, ok := s.symbolPositions[positionKey]

	return ok
}
