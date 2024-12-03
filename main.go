package main

import (
	"fmt"
)

func main() {

	list1, list2 := dealWithInput()

	fmt.Println(distance(list1, list2))
	fmt.Println(similarity(list1, list2))
}
