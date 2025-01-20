package raindrops

import "fmt"

/*
* Problem name: Rain Drops
* Difficult: Easy
 */

func Convert(number int) string {

	var s string

	if number%3 == 0 {
		s += fmt.Sprint("Pling")
	}
	if number%5 == 0 {
		s += fmt.Sprint("Plang")
	}
	if number%7 == 0 {
		s += fmt.Sprint("Plong")
	}

	if s != "" {
		return s
	}

	return fmt.Sprint(number)
}
