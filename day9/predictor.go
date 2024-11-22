package day9

type Predictor struct{}

func (p Predictor) Predict(report Report) int {
	sum := 0

	for i := 0; i < len(report); i++ {
		res := p.predictHistory(report[i])
		sum += res
	}

	return sum
}

func (Predictor) predictHistory(history History) int {
	pastData := make([][]int, 0)
	pastData = append(pastData, history)
	for {
		dataPoint := make([]int, 0)
		for i := 1; i < len(pastData[len(pastData)-1]); i++ {
			dataPoint = append(dataPoint, pastData[len(pastData)-1][i]-pastData[len(pastData)-1][i-1])
		}
		pastData = append(pastData, dataPoint)

		foundLast := true
		for _, d := range dataPoint {
			if d != dataPoint[0] {
				foundLast = false
				break
			}
		}
		if foundLast {
			break
		}
	}
	prediction := 0
	for i := len(pastData) - 1; i >= 0; i-- {
		prediction = pastData[i][0] - prediction
	}

	return prediction
}
