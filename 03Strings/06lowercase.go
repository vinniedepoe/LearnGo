package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	s := strings.ToLower(msg)

	fmt.Println(s)

}
