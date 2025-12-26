package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello")
	fmt.Println("Input")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	fmt.Println("Done")

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		numRating += 1
		fmt.Println(numRating)
	}
}
