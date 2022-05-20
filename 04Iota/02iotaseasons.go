package main

import "fmt"

func main() {
	const (
		Spring = (iota + 1) * 3
		Summer = (iota + 1) * 3
		Fall   = (iota + 1) * 3
		Winter = (iota + 1) * 3
	)

	fmt.Println(Spring, Summer, Fall, Winter)
}
