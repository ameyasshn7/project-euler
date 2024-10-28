package main

import (
	"fmt"
)

func Sieve(N int) int {
	Prime := make([]bool, N+1)
	for i := range Prime {
		Prime[i] = true
	}

	count := 0
	for i := 2; i <= N; i++ {
		if Prime[i] {
			count++
			// Stop when we reach the 10001st prime
			if count == 10001 {
				return i
			}
			for j := i * 2; j <= N; j += i {
				Prime[j] = false
			}
		}
	}
	return -1 // If 10001st prime not found in range
}

func main() {
	N := 200000 // Ensure this is large enough to contain the 10001st prime
	prime10001 := Sieve(N)
	fmt.Printf("The 10001st prime number is %d.\n", prime10001)
}
