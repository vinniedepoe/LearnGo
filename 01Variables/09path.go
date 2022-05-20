package main

import (
	"fmt"
	"path"
)

func main() {
	var dir string

	dir, _ = path.Split("secret/file.txt")

	fmt.Println(dir)

}
