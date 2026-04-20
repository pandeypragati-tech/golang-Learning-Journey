package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "go lang"
	result := strings.ReplaceAll(str, " ", "")
	fmt.Println(result)
}
