package day5_test

import (
	"advent2023/day5"
	"fmt"
	"testing"
)

func TestIsSeed(t *testing.T) {
	tests := []struct {
		seeds     []int
		candidate int
		want      bool
	}{
		{[]int{0, 2, 79, 14}, 92, true},
		{[]int{0, 2, 79, 14}, 93, false},
		{[]int{0, 2, 79, 14}, 79, true},
		{[]int{0, 2, 79, 14}, 78, false},
		{[]int{0, 2, 79, 14}, 0, true},
		{[]int{0, 2, 79, 14}, 2, false},
	}
	almanac := day5.NewAlmanac()
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d IsSeed(%d)", i, tt.candidate),
			func(t *testing.T) {
				almanac.Seeds = []int{0, 2, 79, 14}
				mapper := day5.LocationMapper{
					Almanac: almanac,
				}
				if mapper.IsSeed(tt.candidate) != tt.want {
					t.Errorf("IsSeed(%d) should be %t", tt.candidate, tt.want)
				}
			},
		)
	}
}

func TestGetSeedCandidate(t *testing.T) {
	defaultAlmanac := &day5.Almanac{
		Maps: []day5.AlmanacMap{
			{
				day5.AlmanacMapEntry{0, 69, 1},
				day5.AlmanacMapEntry{1, 0, 69},
			},
			{
				day5.AlmanacMapEntry{60, 56, 37},
				day5.AlmanacMapEntry{56, 93, 4},
			},
		},
	}
	waterToSoil := &day5.Almanac{
		Maps: []day5.AlmanacMap{
			{
				day5.AlmanacMapEntry{0, 15, 37},
				day5.AlmanacMapEntry{37, 52, 2},
				day5.AlmanacMapEntry{39, 0, 15},
			},
			{
				day5.AlmanacMapEntry{49, 53, 8},
				day5.AlmanacMapEntry{0, 11, 42},
				day5.AlmanacMapEntry{42, 0, 7},
				day5.AlmanacMapEntry{57, 7, 4},
			},
		},
	}

	tests := []struct {
		almanac         *day5.Almanac
		candidate, want int
	}{
		{defaultAlmanac, 35, 34},
		{defaultAlmanac, 86, 82},
		{defaultAlmanac, 43, 42},
		{defaultAlmanac, 82, 78},
		{waterToSoil, 81, 81},
		{waterToSoil, 49, 14},
		{waterToSoil, 53, 57},
		{waterToSoil, 41, 13},
		{waterToSoil, 99, 99},

		{defaultAlmanac, 97, 97},
	}
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d: get seed candidate %d", i, tt.candidate),
			func(t *testing.T) {
				mapper := day5.LocationMapper{
					Almanac: tt.almanac,
				}
				if mapper.GetSeedCandidate(tt.candidate) != tt.want {
					t.Errorf("GetSeedCandidate(%d) should be %d", tt.candidate, tt.want)
				}
			})
	}
}
