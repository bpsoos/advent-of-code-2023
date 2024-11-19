package day7_test

import (
	"advent2023/day7"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseGamesWithSingle(t *testing.T) {
	tests := []struct {
		rawGames string
		want     []day7.Game
	}{
		{
			"32T3K 765",
			day7.Games{{day7.Hand{day7.Three, day7.Two, day7.Ten, day7.Three, day7.King}, 765}},
		},
		{
			"AKQJ9 28 ",
			day7.Games{{day7.Hand{day7.Ace, day7.King, day7.Queen, day7.Jack, day7.Nine}, 28}},
		},
	}

	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d parse game", i),
			func(t *testing.T) {
				parsed := day7.ParseGames(strings.NewReader(tt.rawGames))
				if len(parsed) != len(tt.want) {
					t.Errorf("parsed games len don't match expected\nparsed:  %v\nexpected: %v", len(parsed), len(tt.want))
				}
				for i, game := range tt.want {
					if !reflect.DeepEqual(parsed[i].Hand, game.Hand) {
						t.Errorf("parsed hand doesn't match expected\nparsed:  %v\nexpected: %v", parsed[i].Hand, game.Hand)
					}
					if !reflect.DeepEqual(parsed[i].Bid, game.Bid) {
						t.Errorf("parsed bid doesn't match expected\nparsed:  %v\nexpected: %v", parsed[i].Bid, game.Bid)
					}
				}
			},
		)
	}
}
