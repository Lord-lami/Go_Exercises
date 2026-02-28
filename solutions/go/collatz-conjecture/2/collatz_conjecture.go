// Package collatzconjecture provides a function for counting collatz countdown steps
package collatzconjecture

import "errors"

// CollatzConjecture takes a number n and calculates the amount of steps
// it takes to to reach 1 with a collatz countdown
// It returns an error if n is less than 1.
func CollatzConjecture(n int) (steps int, err error) {
	if n < 1 {
		err = errors.New("stubid! only positive integers allowed")
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return
}
