package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var list = []string{"a", "aa", "aaa"}
	fmt.Printf("type: %T \n", list)

	list = append(list, "aaaa", "5a")
	fmt.Println(list)

	list = append(list[1:]) //Start from 1st index
	fmt.Println(list)

	list = append(list[1:3]) //Only keep index 1(include) to 3(exclude)
	fmt.Println(list)

	list = append(list[:3]) //Only keep to 3(exclude)
	fmt.Println(list)

	// Make with slices
	scores := make([]int, 4) // Returns slice ?
	scores[0] = 3
	scores[1] = 2
	scores[2] = 1
	scores[3] = 0
	scores = append(scores, 4, 5, 6) // reallocates memory
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println("Sorted", scores)

	// remove value from slice based on index
	var newList = []string{"0", "1", "2", "3", "4", "5"}
	fmt.Println(newList)
	index := 2
	newList = append(newList[:index], newList[index+1:]...)
	fmt.Println(newList)
}
