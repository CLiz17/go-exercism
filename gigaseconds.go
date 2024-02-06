// Package gigasecond is to determine the date and time one gigasecond after a certain date.
package gigasecond

import "time"

// AddGigasecond function is to add one gigasecond after a certain date.
func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * 1000000000)
}
