package day5

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

func (lm LocationMapper) getEndLocation(seed int) int {
	location := seed
	for i := 0; i < len(lm.Almanac.Maps); i++ {
		for j := 0; j < len(lm.Almanac.Maps[i]); j++ {
			newLocation := updateLocation(location, lm.Almanac.Maps[i][j])
			if newLocation != location {
				location = newLocation
				break
			}
		}
	}

	return location
}

func updateLocation(location int, mapEntry AlmanacMapEntry) int {
	if location < mapEntry.SrcRangeStart() {
		return location
	}
	if location > mapEntry.SrcRangeEnd() {
		return location
	}

	location += mapEntry.Offset()

	return location
}
