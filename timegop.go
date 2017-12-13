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
	"fmt"
	"time"
)

type mode int

const (
	natural mode = iota
	freezing
	traveling
)

var currentMode mode = natural
var traveledTime time.Time

// Freeze fixes current time as given t
func Freeze(t time.Time) func() {
	currentMode = freezing
	traveledTime = t
	return Return
}

func Travel(t time.Time) func() {
	return Return
}

// Now returns a fixed time if Freeze is called, otherwise a real time
func Now() time.Time {
	switch currentMode {
	case natural:
		return time.Now()
	case freezing:
		return traveledTime
	}
	panic(fmt.Sprintf("undefined mode %d", currentMode))
}

// Return disable time fixing
func Return() {
	currentMode = natural
}

// Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).
func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}

// Until returns the duration until t. It is shorthand for t.Sub(time.Now())
func Until(t time.Time) time.Duration {
	return t.Sub(Now())
}
