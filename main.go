package main

import (
	"fmt"
)

func main() {

	// Day 1
	list1, list2 := dealWithInput()

	fmt.Println(distance(list1, list2))
	fmt.Println(similarity(list1, list2))

	// Day 2
	reports := consumeData()

	fmt.Println(reports)
}
