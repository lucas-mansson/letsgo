package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://lundakarnevalen.se"

func main() {
	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("%T\n", response)
	fmt.Println(response)

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
