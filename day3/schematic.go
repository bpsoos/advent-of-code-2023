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
		Position PartPosition
	}

	PartPosition mathx.LineSegment
)

func (p PartNumber) AdjPositions() []mathx.Point {
	return p.Position.AdjPositions()
}

func (pp PartPosition) AdjPositions() []mathx.Point {
	positionCount := 6 + pp.GetLength()*2
	adjPositions := make([]mathx.Point, 0, positionCount)

	adjPositions = append(adjPositions, mathx.Point{X: pp.Start.X - 1, Y: pp.Start.Y - 1})
	adjPositions = append(adjPositions, mathx.Point{X: pp.Start.X - 1, Y: pp.Start.Y})
	adjPositions = append(adjPositions, mathx.Point{X: pp.Start.X - 1, Y: pp.Start.Y + 1})

	adjPositions = append(adjPositions, mathx.Point{X: pp.End.X + 1, Y: pp.End.Y - 1})
	adjPositions = append(adjPositions, mathx.Point{X: pp.End.X + 1, Y: pp.End.Y})
	adjPositions = append(adjPositions, mathx.Point{X: pp.End.X + 1, Y: pp.End.Y + 1})

	for i := 0; i < pp.GetLength(); i++ {
		adjPositions = append(adjPositions, mathx.Point{X: pp.Start.X + i, Y: pp.Start.Y - 1})
		adjPositions = append(adjPositions, mathx.Point{X: pp.Start.X + i, Y: pp.Start.Y + 1})
	}

	return adjPositions
}
func (pp PartPosition) GetLength() int {
	return pp.End.X - pp.Start.X + 1
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
