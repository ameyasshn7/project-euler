/*Problem 1:
If we list all the natural numbers below $10$ that are multiples of $3$ or $5$, we get $3, 5, 6$ and $9$. The sum of these multiples is $23$.
Find the sum of all the multiples of $3$ or $5$ below $1000$.
*/

package main

import "fmt"

func isMultiple(number int, bases []int) bool {
	for _, base := range bases {
		if number%base == 0 {
			return true
		}
	}
	return false
}

func main() {
	number := 1000
	sum := 0
	bases := []int{3, 5}

	for i := 0; i < number; i++ {
		if isMultiple(i, bases) {
			sum += i
		}
	}
	fmt.Println(sum)
}
