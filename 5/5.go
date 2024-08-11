//

package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func smallestMultiple(limit int) int {
	result := 1

	for i := 2; i < limit; i++ {
		result = lcm(result, i)
	}
	return result
}
func main() {
	fmt.Println("Smallest number that is evenly divisble by all numbers in the rage 1 to 20: ", smallestMultiple((20)))
}
