package main

import (
	"fmt"
	"math"
	"math/big"
)

// Naive approach for finding the largest prime factor of an int64 number
func largestPrimeFactorNaive(n int64) int64 {
	largestPrime := int64(1)
	for n%2 == 0 {
		largestPrime = 2
		n /= 2
	}

	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			largestPrime = i
			n /= i
		}
	}

	if n > 2 {
		largestPrime = n
	}

	return largestPrime
}

// Pollard's Rho algorithm to find a factor of a large number n
func PollardsRho(n *big.Int) *big.Int {
	x := big.NewInt(2)
	y := big.NewInt(2)
	d := big.NewInt(1)
	one := big.NewInt(1)
	c := big.NewInt(1) // Constant for f(x) = (x^2 + c) % n. (in this case c = 1)

	for d.Cmp(one) == 0 {
		// f(x) = (x^2 + c) % n
		// x = f(x)
		// tortoise
		x.Mul(x, x)
		x.Add(x, c)
		x.Mod(x, n)

		// y = f(f(y))
		// hare
		y.Mul(y, y)
		y.Add(y, c)
		y.Mod(y, n)
		y.Mul(y, y)
		y.Add(y, c)
		y.Mod(y, n)

		d.Sub(x, y)
		d.Abs(d)
		d.GCD(nil, nil, d, n)
	}

	if d.Cmp(n) == 0 {
		return nil
	}

	return d
}

// Function to find the largest prime factor using Pollard's Rho and fallback to naive
func findLargest(n *big.Int) *big.Int {
	largest := big.NewInt(1)

	for n.Cmp(big.NewInt(1)) > 0 {
		if n.ProbablyPrime(20) {
			if n.Cmp(largest) > 0 {
				largest.Set(n)
			}
			break
		}

		factor := PollardsRho(n)
		if factor == nil {
			// Fallback to the naive method for remaining n
			fmt.Println("Pollards Rho Failed")
			largestPrime := largestPrimeFactorNaive(n.Int64())
			return big.NewInt(largestPrime)
		}

		for new(big.Int).Mod(n, factor).Cmp(big.NewInt(0)) == 0 {
			n.Div(n, factor)
		}

		if factor.Cmp(largest) > 0 {
			largest.Set(factor)
		}
	}

	return largest
}

func main() {
	number := big.NewInt(600851475143)
	largestPrime := findLargest(number)
	fmt.Println("The largest prime factor is:", largestPrime)
}
