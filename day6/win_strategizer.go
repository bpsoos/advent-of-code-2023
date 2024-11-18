package day6

type WinStrategizer struct {
	TimeSheet []RaceTime
}

func NewWinStrategizer(timeSheet []RaceTime) *WinStrategizer {
	return &WinStrategizer{TimeSheet: timeSheet}
}

func (ws WinStrategizer) WinPossibilities() []int {
	possibilities := make([]int, len(ws.TimeSheet))

	for i := 0; i < len(ws.TimeSheet); i++ {
		possibilities[i] = ws.getPossibilityCount(ws.TimeSheet[i])
	}

	return possibilities
}

func (ws WinStrategizer) getPossibilityCount(race RaceTime) int {
	possibilityCount := 0

	for holdTime := range race.Time {
		runTime := race.Time - holdTime
		distance := runTime * holdTime
		if distance > race.Distance {
			possibilityCount++
		}
	}

	return possibilityCount
}
