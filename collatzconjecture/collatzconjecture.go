package collatzconjecture

import "fmt"

/*
* Problem Name: Collatz Conjecture
* Difficulty: Easy
 */

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, fmt.Errorf("Number {n} must be in the set N*, got: %v", n)
	}

	if n == 1 {
		return 0, nil
	}

	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n /= 2
			steps++
		} else {
			n = (3 * n) + 1
			steps++
		}
	}

	return steps, nil
}
