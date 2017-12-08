/*
  Copyright (C) 2017 Akira Tanimura (@AuToPP)

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

// Package timegop provides stubing time.Now
package timegop

import (
	"time"
)

var frozen = false
var freezedTime time.Time

// Freeze fixes current time as given t
func Freeze(t time.Time) func() {
	frozen = true
	freezedTime = t

	return Return
}

// Now returns a fixed time if Freeze is called, otherwise a real time
func Now() time.Time {
	if !frozen {
		return time.Now()
	}
	return freezedTime
}

// Return disable time fixing
func Return() {
	frozen = false
}
