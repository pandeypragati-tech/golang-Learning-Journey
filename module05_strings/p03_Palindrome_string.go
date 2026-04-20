package main

import "fmt"

func main() {
	str := "madam"
	rev := ""

	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}

	if str == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
