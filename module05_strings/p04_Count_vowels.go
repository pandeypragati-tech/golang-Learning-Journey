package main

import "fmt"

func main() {
	str := "hello world"
	count := 0

	for _, ch := range str {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}
	fmt.Println("Vowels:", count)
}
