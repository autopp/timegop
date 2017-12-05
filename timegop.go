package timegop

import (
	"time"
)

var frozen = false
var freezedTime time.Time

func Freeze(t time.Time) {
	frozen = true
	freezedTime = t
}

func Now() time.Time {
	if frozen {
		return freezedTime
	} else {
		return time.Now()
	}
}

func Return() {
	frozen = false
}
