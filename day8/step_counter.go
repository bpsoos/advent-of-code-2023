package day8

import (
	"advent2023/btree"
	"advent2023/mathx"
	"strings"
)

type StepCounter struct{}

func (sc StepCounter) Count(stepPattern string, network *btree.Tree[string]) int {
	return sc.count(stepPattern, network, func(val string) bool { return val == "ZZZ" })
}
func (sc StepCounter) count(stepPattern string, network *btree.Tree[string], compareFunc func(string) bool) int {
	if network == nil {
		return 0
	}

	count := 0
	stepIndex := 0

	node := network
	for {
		count++
		if string(stepPattern[stepIndex]) == "R" {
			node = node.Right
		}
		if string(stepPattern[stepIndex]) == "L" {
			node = node.Left
		}
		if compareFunc(node.Value) {
			break
		}

		stepIndex++
		if stepIndex >= len(stepPattern) {
			stepIndex = 0
		}
	}

	return count
}

func (sc StepCounter) CountAsGhost(stepPattern string, network map[string]*btree.Tree[string]) int {
	nodes := make([]*btree.Tree[string], 0)
	cycleSteps := make([]int, 0)

	for k, v := range network {
		if strings.Contains(k, "A") {
			nodes = append(nodes, v)
			cycleSteps = append(cycleSteps,
				sc.count(
					stepPattern,
					v,
					func(val string) bool { return strings.Contains(val, "Z") },
				),
			)
		}
	}

	return mathx.LCM(cycleSteps[0], cycleSteps[1], cycleSteps[2:]...)
}
