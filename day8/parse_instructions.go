package day8

import (
	"advent2023/btree"
	"bufio"
	"io"
	"strings"
)

func ParseInstructions(rawNetwork io.Reader) (string, map[string]*btree.Tree[string]) {
	scanner := bufio.NewScanner(rawNetwork)
	scanner.Scan()
	stepPattern := scanner.Text()
	return stepPattern, parseNetwork(scanner)
}

func parseNetwork(scanner *bufio.Scanner) map[string]*btree.Tree[string] {
	treeMap := make(map[string]*btree.Tree[string])
	networkMap := make(map[string][2]string)
	scanner.Scan()

	for scanner.Scan() {
		key, values := parseNode(scanner.Text())
		networkMap[key] = values
		node := new(btree.Tree[string])
		node.Value = key
		treeMap[key] = node
	}

	for k, v := range networkMap {
		current := treeMap[k]

		left := treeMap[v[0]]
		if left != current {
			current.Left = left
		}

		right := treeMap[v[1]]
		if right != current {
			current.Right = right
		}
	}

	return treeMap
}

func parseNode(line string) (string, [2]string) {
	splitLine := strings.Split(line, " = ")
	key := splitLine[0]
	values := strings.Split(strings.Trim(splitLine[1], "()"), ", ")
	if len(values) != 2 {
		panic("invalid values")
	}

	return key, [2]string{values[0], values[1]}
}
