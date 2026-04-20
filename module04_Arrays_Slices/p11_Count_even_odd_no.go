package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	even, odd := 0, 0

	for _, v := range arr {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("Even:", even, "Odd:", odd)
}
