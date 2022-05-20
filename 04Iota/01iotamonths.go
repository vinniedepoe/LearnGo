package main

import "fmt"

func main() {
	const (
		Jan = iota + 1
		Feb = iota + 1
		Mar = iota + 1
	)

	fmt.Println(Jan, Feb, Mar)
}
