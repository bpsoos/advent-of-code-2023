package day8

import (
	"advent2023/btree"
	"fmt"
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
	stepIndex := 0
	nodes := make([]*btree.Tree[string], 0)

	for k, v := range network {
		if strings.Contains(k, "A") {
			nodes = append(nodes, v)
		}
	}

	ghostCount := len(nodes)
	count := 0
	for {
		count++
		newNodes := make([]*btree.Tree[string], len(nodes))
		for i := 0; i < ghostCount; i++ {
			if string(stepPattern[stepIndex]) == "R" {
				newNodes[i] = nodes[i].Right
			}
			if string(stepPattern[stepIndex]) == "L" {
				newNodes[i] = nodes[i].Left
			}
		}
		nodes = newNodes

		found := true
		for i := 0; i < ghostCount; i++ {
			fmt.Printf("%s ", nodes[i].Value)
			if !strings.Contains(nodes[i].Value, "Z") {
				found = false
				fmt.Print(" NO ")
			}
		}
		fmt.Print("\n")
		if found {
			break
		}

		stepIndex++
		if stepIndex >= len(stepPattern) {
			stepIndex = 0
		}
	}

	return count
}
