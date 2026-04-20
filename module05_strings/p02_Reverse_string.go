package main

import "fmt"

func main() {
	str := "hello"
	rev := ""

	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	fmt.Println(rev)
}
