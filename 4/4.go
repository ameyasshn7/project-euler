/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is $9009 = 91 \times 99$.
Find the largest palindrome made from the product of two $3$-digit numbers.
*/

package main

import "fmt"

//Find the largest Palindrome within any given number range

func makeNumber(num1 int, num2 int) int {
	largest_palindrome := 0
	for i := num2; i >= num1; i-- {

		if i*num2 < largest_palindrome {
			break
		}

		for j := i; j >= num1; j-- {
			product := i * j

			if product < largest_palindrome {
				break
			}

			if isPalindrome(product) {
				largest_palindrome = product
			}
		}

	}
	return largest_palindrome
}

//check if a number is a palindrome
func isPalindrome(number int) bool {
	original := number
	rev := 0
	for number > 0 {
		rem := number % 10
		rev = rev*10 + rem
		number /= 10
	}
	return rev == original
}
func main() {
	fmt.Println("The largest 3 Digit palindrome number is:", makeNumber(100, 999))

}
