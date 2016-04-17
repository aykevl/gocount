// Copyright 2016 Ayke van Laethem.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

// Package gocount counts the number of goroutines currently running.
// It should only be used for testing purposes, as it is implemented in a very
// inefficient way.
package gocount

import (
	"bytes"
	"regexp"
	"runtime/pprof"
)

var routineRegexp = regexp.MustCompile("(^|\n\n)goroutine ")

// Number returns the number of goroutines, not including runtime goroutines
// like those created by the garbage collector.
//
// Also take a look at runtime.NumGoroutine(), which counts all goroutines
// including those created by the runtime.
func Number() int {
	profile := pprof.Lookup("goroutine")
	buf := &bytes.Buffer{}
	profile.WriteTo(buf, 2)
	return len(routineRegexp.FindAll(buf.Bytes(), -1))
}
