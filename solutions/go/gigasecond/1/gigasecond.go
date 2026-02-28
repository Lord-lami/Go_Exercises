// Package gigasecond provides tools for measuring time only in seconds.
package gigasecond

import "time"

// AddGigasecond takes t, a time.Time variable, and 
// returns the time.Time after 1 Gigasecond (1 billion seconds) have passed.
func AddGigasecond(t time.Time) time.Time {
	gigaSecond, _ := time.ParseDuration("1000000000s")
	return t.Add(gigaSecond)
}
