package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("List of all", languages)
	fmt.Println("JS", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of all", languages)

	for key := range languages {
		fmt.Println(languages[key])
	}

	volvo := &Car{"volvo", 1, false}
	fmt.Println(volvo)
	fmt.Printf("%T\n", volvo)
}

type Car struct {
	Name string
	Year int
	A    bool
}
