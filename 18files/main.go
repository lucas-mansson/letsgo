package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "In to the file"

	file, err := os.Create("./file.txt")
	checkNilErr(err)

	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println(length)
	readFile("./file.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
