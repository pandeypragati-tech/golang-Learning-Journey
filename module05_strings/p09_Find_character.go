package main

import "fmt"

func main() {
	str := "golang"
	ch := 'a'
	found := false

	for _, c := range str {
		if c == ch {
			found = true
			break
		}
	}

	fmt.Println("Found:", found)
}
