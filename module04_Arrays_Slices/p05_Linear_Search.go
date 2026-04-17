package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8}
	key := 6
	found := false

	for _, v := range arr {
		if v == key {
			found = true
			break
		}
	}

	fmt.Println("Found:", found)
}
