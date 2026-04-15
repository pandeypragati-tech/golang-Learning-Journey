package main

import "fmt"

func main() {
	var num int
	isPrime := true

	fmt.Println("Enter the Number:")
	fmt.Scanln(&num)

	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Println("Number is Prime", num)
	} else {
		fmt.Println("Number is Not Prime", num)
	}

}
