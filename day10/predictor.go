package day10

import "fmt"

type Predictor struct{}

func (p Predictor) Predict(report Report) int {
	sum := 0

	for i := 0; i < len(report); i++ {
		fmt.Println("-----------------")
		res := p.predictHistory(report[i])
		sum += res
	}

	return sum
}

func (Predictor) predictHistory(history History) int {
	pastData := make([][]int, 0)
	pastData = append(pastData, history)
	fmt.Println(history)
	for {
		dataPoint := make([]int, 0)
		for i := 1; i < len(pastData[len(pastData)-1]); i++ {
			dataPoint = append(dataPoint, pastData[len(pastData)-1][i]-pastData[len(pastData)-1][i-1])
		}
		pastData = append(pastData, dataPoint)

		foundLast := true
		fmt.Println(dataPoint)
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
		fmt.Println("-----")
		fmt.Printf("%d, %v", prediction, pastData[i])
	}
	fmt.Println("-----")

	fmt.Println(prediction)
	return prediction
}
