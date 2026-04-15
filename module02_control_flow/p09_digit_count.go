package main

import "fmt"

func main() {
	var num, count int

	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	for num != 0 {
		num /= 10
		count++
	}

	fmt.Println("Digits:", count)
}
