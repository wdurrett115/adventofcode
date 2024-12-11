package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func validate(direction bool, report []int) (int, int) {
	for i := 0; i < len(report)-1; i++ {

		diff := math.Abs(float64(report[i])) - float64(report[i+1])

		if !(diff >= 1 && diff <= 3) {
			return 2, i
		}

		if !direction {
			if report[i] >= report[i+1] {
				return 0, i
			}
		} else if direction {
			if report[i] <= report[i+1] {
				return 0, i
			}
		}
	}

	return 1, 0
}

func safety(reports [][]int) (int, [][]int, int) {
	var amount int
	var goods [][]int
	var direction bool
	var safe int
	var failedIndex int
	for _, report := range reports {
		if report[0] < report[1] {
			direction = false
		} else if report[0] == report[1] {
			continue
		} else {
			direction = true
		}

		safe, failedIndex = validate(direction, report)
		if safe == 1 {
			goods = append(goods, report)
		}
		amount += safe
	}
	return amount, goods, failedIndex
}

func consumeData() [][]int {
	data := `1 3 6 7 9`

	var reports [][]int
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		fields := strings.Fields(line)
		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
			}
			row = append(row, num)
		}

		reports = append(reports, row)
	}

	return reports
}
