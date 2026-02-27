package collatzconjecture

import "errors"

func CollatzConjecture(n int) (steps int, err error) {
	if n == 1 {
		return
	}
	if n < 1 {
		err = errors.New("Only Positive integers are allowed")
		return
	}
	if n%2 == 0 {
		steps, err = CollatzConjecture(n / 2)
		steps++
	} else {
		steps, err = CollatzConjecture(n*3 + 1)
		steps++
	}
	return
}
