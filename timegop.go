package timegop

import (
	"time"
)

var frozen = false
var freezedTime time.Time

// Freeze fixes current time as given t
func Freeze(t time.Time) {
	frozen = true
	freezedTime = t
}

// Now returns a fixed time if Freeze is called, otherwise a real time
func Now() time.Time {
	if frozen {
		return freezedTime
	}
	return time.Now()
}

// Return disable time fixing
func Return() {
	frozen = false
}
