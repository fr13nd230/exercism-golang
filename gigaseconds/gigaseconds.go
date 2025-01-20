package gigaseconds

import "time"

/*
* Problem's Name: Gigaseconds
* Difficulty: Easy
 */

func AddGigasecond(t time.Time) time.Time {
	g := 1_000_000_000
	finalTime := int(time.Second) * g
	return t.Add(time.Duration(finalTime))
}
