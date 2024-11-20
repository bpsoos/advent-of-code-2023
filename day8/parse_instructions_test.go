package day8_test

import (
	"advent2023/btree"
	"advent2023/day8"
	"strings"
	"testing"
)

func TestParseSimpleInstruction(t *testing.T) {
	instructions := strings.NewReader(strings.TrimSpace(
		`RL

AAA = (BBB, CCC)
BBB = (CCC, CCC)
CCC = (CCC, CCC)
`))
	wantSteps := "RL"
	c := btree.Tree[string]{
		Right: nil,
		Left:  nil,
		Value: "CCC",
	}
	b := btree.Tree[string]{
		Right: &c,
		Left:  &c,
		Value: "BBB",
	}
	a := btree.Tree[string]{
		Right: &c,
		Left:  &b,
		Value: "AAA",
	}
	wantNetwork := a
	steps, network := day8.ParseInstructions(instructions)
	if network.String() != wantNetwork.String() {
		t.Errorf("parsed network doesn't match want\nparsed:  %v\nexpected: %v\n", network, wantNetwork)
	}
	if steps != wantSteps {
		t.Errorf("parsed steps doesn't match want\nparsed:  %v\nexpected: %v\n", steps, wantSteps)
	}

}
