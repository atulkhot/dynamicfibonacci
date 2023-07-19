package main

import (
	"fmt"
	"strconv"
)

var fibonacci_values []int64

func fibonacci_on_the_fly(n int64) int64 {
	if int64(len(fibonacci_values)) > n {
		return fibonacci_values[n]
	}
	result := fibonacci_on_the_fly(n-1) + fibonacci_on_the_fly(n-2)
	fibonacci_values = append(fibonacci_values, result)
	return result
}

func initialize_slice() {
	fibonacci_values = make([]int64, 93)
	fibonacci_values[0] = 0
	fibonacci_values[1] = 1

	for i := 2; i < len(fibonacci_values); i++ {
		fibonacci_values[i] = fibonacci_values[i-1] + fibonacci_values[i-2]
	}
}

func fibonacci_prefilled(n int64) int64 {
	if int(n) >= len(fibonacci_values) {
		panic("Value given is out of range")
	}
	return fibonacci_values[n]
}

func fibonacci_bottom_up(n int64) int64 {
	if n <= 1 {
		return int64(n)
	}

	var fib_i, fib_i_minus_1, fib_i_minus_2 int64
	fib_i_minus_2 = 0
	fib_i_minus_1 = 1
	fib_i = fib_i_minus_1 + fib_i_minus_2
	for i := int64(1); i < n; i++ {
		// Calculate this value of fib_i.
		fib_i = fib_i_minus_1 + fib_i_minus_2

		// Set fib_i_minus_2 and fib_i_minus_1 for the next value.
		fib_i_minus_2 = fib_i_minus_1
		fib_i_minus_1 = fib_i
	}
	return fib_i
}

func main() {
	// Fill-on-the-fly.
	fibonacci_values = make([]int64, 2)
	fibonacci_values[0] = 0
	fibonacci_values[1] = 1

	// Prefilled.
	initialize_slice()

	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)

		// Uncomment one of the following.
		//fmt.Printf("fibonacci_on_the_fly(%d) = %d\n", n, fibonacci_on_the_fly(n))
		//fmt.Printf("fibonacci_prefilled(%d) = %d\n", n, fibonacci_prefilled(n))
		fmt.Printf("fibonacci_bottom_up(%d) = %d\n", n, fibonacci_bottom_up(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacci_values); i++ {
		fmt.Printf("%d: %d\n", i, fibonacci_values[i])
	}
}
