package main

import "fmt"

func isPalindrome(num int) bool {
	rev := 0
	temp := num

	for num != 0 {
		rev = rev*10 + num%10
		num /= 10
	}
	return temp == rev
}

func main() {
	fmt.Println(isPalindrome(121))
}
