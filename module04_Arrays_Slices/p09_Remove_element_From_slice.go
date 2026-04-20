package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	index := 1

	s = append(s[:index], s[index+1:]...)

	fmt.Println(s)
}
