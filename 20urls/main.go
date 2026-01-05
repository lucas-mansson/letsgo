package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://example.com:3000/learn?coursename=reactjs&?id=hello"

func main() {
	fmt.Println(myurl)
	res, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("%T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["?id"])
	for key, val := range qparams {
		fmt.Println(key)
		fmt.Println(val)
	}
}
