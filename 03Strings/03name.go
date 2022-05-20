package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	msg := `
	hi ` + name + `
	How are you?`

	fmt.Printf(msg)
}
