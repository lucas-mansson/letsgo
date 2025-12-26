package main

import "fmt"

func main() {
	var list [4]string
	list[0] = "0"
	list[1] = "1"
	list[2] = "2"
	fmt.Println(list)

	var list2 = [3]string{"0", "1", "2"}
	fmt.Println(list2)
	fmt.Println(len(list2))
}
