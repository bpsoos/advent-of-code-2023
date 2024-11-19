package day7_test

import (
	"advent2023/day7"
	"fmt"
	"testing"
)

func TestHandTypes(t *testing.T) {
	tests := []struct {
		hand day7.Hand
		want day7.HandType
	}{
		{
			day7.Hand{
				day7.Two,
				day7.Three,
				day7.Four,
				day7.Five,
				day7.Six,
			},
			day7.HighCard,
		},
		{
			day7.Hand{
				day7.Two,
				day7.Two,
				day7.Two,
				day7.Two,
				day7.Two,
			},
			day7.FiveOfAKind,
		},
	}
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d hand types", i),
			func(t *testing.T) {
				handType := tt.hand.Type()
				if handType != tt.want {
					t.Errorf("hand type doesn't match expected\n type: %v\nexpected: %v\n", handType, tt.want)
				}
			},
		)
	}

}
