package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is fun"
	words := strings.Fields(str)

	fmt.Println("Words:", len(words))
}
