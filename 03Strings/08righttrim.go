package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Marko              "
	trim := strings.TrimRight(name, " ")
	fmt.Println(len(trim))
}
