package main

import "fmt"

func main() {
	var num, rev, temp int

	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	temp = num

	for num != 0 {
		rev = rev*10 + num%10
		num /= 10
	}

	if temp == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
