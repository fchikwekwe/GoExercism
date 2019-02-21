// Package gigasecond calulates the time a gigasecond after the given input
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigasecond of time to the given input
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
