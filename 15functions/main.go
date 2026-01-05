package main

import "fmt"

func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println(adder(1, 1, 1, 1))
	u := &User{"a", "b", false, 1, 2}
	fmt.Printf("%+v\n", u)
	fmt.Println(u.GetStatus())
	u.ChangeEmail()
	fmt.Printf("%+v\n", u)
}

// function
func adder(values ...int) (int, string) {
	// anonymous function
	func() {
		fmt.Println("Hello")
	}()

	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hello"
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() bool {
	return u.Status
}

func (u *User) ChangeEmail() {
	u.Email = "newEmail"
}
