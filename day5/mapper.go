package day5

import (
	"advent2023/mathx"
)

type LocationMapper struct {
	Almanac *Almanac
}

func (lm LocationMapper) Locations() []int {
	seedCount := len(lm.Almanac.Seeds)
	endLocations := make([]int, seedCount)

	for i := 0; i < seedCount; i++ {
		endLocations[i] = lm.getEndLocation(lm.Almanac.Seeds[i])
	}

	return endLocations
}

func (lm LocationMapper) RangedSeedLocations() int {
	location := 0
	for {
		if lm.IsSeed(lm.GetSeedCandidate(location)) {
			return location
		}
		location++
	}
}

func (lm LocationMapper) IsSeed(candidate int) bool {
	for i := 0; i < len(lm.Almanac.Seeds); i += 2 {
		if candidate >= lm.Almanac.Seeds[i] && candidate < lm.Almanac.Seeds[i]+lm.Almanac.Seeds[i+1] {
			return true
		}
	}

	return false
}

func (lm LocationMapper) GetSeedCandidate(endLocation int) int {
	location := endLocation
	for i := len(lm.Almanac.Maps) - 1; i >= 0; i-- {
		for j := 0; j < len(lm.Almanac.Maps[i]); j++ {
			newLocation := lm.updateLocationBackwards(location, lm.Almanac.Maps[i][j])
			if newLocation != location {
				location = newLocation
				break
			}
		}
	}

	return location
}

func (lm LocationMapper) updateLocationBackwards(location int, mapEntry AlmanacMapEntry) int {
	if location < mapEntry.DstRangeStart() {
		return location
	}
	if location > mapEntry.DstRangeEnd() {
		return location
	}

	return location - mapEntry.Offset()
}

func (lm LocationMapper) getEndLocation(seed int) int {
	location := seed
	for i := 0; i < len(lm.Almanac.Maps); i++ {
		for j := 0; j < len(lm.Almanac.Maps[i]); j++ {
			newLocation := lm.updateLocation(location, lm.Almanac.Maps[i][j])
			if newLocation != location {
				location = newLocation
				break
			}
		}
	}

	return location
}

func (lm LocationMapper) updateLocation(location int, mapEntry AlmanacMapEntry) int {
	if location < mapEntry.SrcRangeStart() {
		return location
	}
	if location > mapEntry.SrcRangeEnd() {
		return location
	}

	location = mathx.AddInt(location, mapEntry.Offset())

	return location
}
