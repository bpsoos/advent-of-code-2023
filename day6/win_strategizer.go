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
		possibilities[i] = ws.possibilitiesForRace(ws.TimeSheet[i])
	}
	return possibilities
}

func (ws WinStrategizer) possibilitiesForRace(race RaceTime) int {
	return 0
}
