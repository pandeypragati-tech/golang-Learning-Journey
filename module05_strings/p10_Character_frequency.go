package main

import "fmt"

func main() {
	str := "hello"
	freq := make(map[rune]int)

	for _, ch := range str {
		freq[ch]++
	}

	fmt.Println(freq)
}
