package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Direct formula approach
func sumOfSquares(limit int) int {
	return (limit * (limit + 1) * (2*limit + 1)) / 6
}

func squareOfSums(limit int) int {
	sum := (limit * (limit + 1)) / 2
	return sum * sum
}

func calculateDifference(limit int) int {
	return int(math.Abs(float64(sumOfSquares(limit) - squareOfSums(limit))))
}

// Memoization approach
var memo map[int]int

func sumOfSquaresMemoized(n int) int {
	if val, found := memo[n]; found {
		return val
	}
	if n == 1 {
		return n
	}
	result := sumOfSquaresMemoized(n-1) + n*n
	memo[n] = result
	return result
}

func calculateDifferenceMemoized(limit int) int {
	return squareOfSums(limit) - sumOfSquaresMemoized(limit)
}

// Iterative approach
func sumOfSquaresIterative(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSumsIterative(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum * sum
}

func calculateDifferenceIterative(limit int) int {
	sumOfSquares := sumOfSquaresIterative(limit)
	squareOfSums := squareOfSumsIterative(limit)
	return squareOfSums - sumOfSquares
}

// Function to measure both time and space complexity
func compareComplexity(limit int) {
	printMemUsage("Before any calculation")

	// Measure time and space for Direct Formula
	start := time.Now()
	calculateDifference(limit)
	directTime := time.Since(start)
	printMemUsage("After Direct Formula")

	// Measure time and space for Memoization
	memo = make(map[int]int)
	start = time.Now()
	calculateDifferenceMemoized(limit)
	memoizedTime := time.Since(start)
	printMemUsage("After Memoization")

	// Measure time and space for Iterative Approach
	start = time.Now()
	calculateDifferenceIterative(limit)
	iterativeTime := time.Since(start)
	printMemUsage("After Iterative Approach")

	// Output the timing results
	fmt.Printf("Time Complexity:\n")
	fmt.Printf("Direct Formula: %v\n", directTime)
	fmt.Printf("Memoized Approach: %v\n", memoizedTime)
	fmt.Printf("Iterative Approach: %v\n", iterativeTime)
}

func printMemUsage(phase string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s - Alloc = %v KB, TotalAlloc = %v KB, Sys = %v KB, NumGC = %v\n", phase, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

func main() {
	var limit int
	fmt.Print("Enter a value for n: ")
	fmt.Scan(&limit)

	compareComplexity(limit)
}
