package timegop

import (
	"testing"
	"time"
)

func TestFreeze(t *testing.T) {
	freezedTime := time.Date(2017, time.December, 6, 9, 29, 0, 0, time.Local)
	Freeze(freezedTime)
	expected := freezedTime.UnixNano()
	actual := Now().UnixNano()

	if expected != actual {
		t.Errorf("actual = %d, expected = %d", actual, expected)
	}
}
