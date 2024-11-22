package day9

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type History []int
type Report []History

func ParseReport(rawReport io.Reader) Report {
	scanner := bufio.NewScanner(rawReport)
	report := make(Report, 0)
	for scanner.Scan() {
		report = append(report, parseLine(scanner.Text()))
	}

	return report
}

func parseLine(line string) History {
	values := strings.Split(line, " ")
	history := make(History, len(values))
	for i := 0; i < len(values); i++ {
		value, err := strconv.Atoi(values[i])
		if err != nil {
			panic(err)
		}
		history[i] = value
	}

	return history
}
