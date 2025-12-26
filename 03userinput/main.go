package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello there"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("Type %T \n", input)
}
