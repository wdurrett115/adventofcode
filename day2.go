package main

import (
	"fmt"
	"strconv"
	"strings"
)

func safety(reports [][]int) int {

	return 0
}

func consumeData() [][]int {
	data := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

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
