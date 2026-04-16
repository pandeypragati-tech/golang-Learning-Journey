package main

import "fmt"

func isArmstrong(num int) bool {
	temp := num
	sum := 0

	for temp != 0 {
		digit := temp % 10
		sum += digit * digit * digit
		temp /= 10
	}

	return sum == num
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if isArmstrong(num) {
		fmt.Println("Armstrong Number")
	} else {
		fmt.Println("Not Armstrong Number")
	}
}
