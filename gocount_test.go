// Copyright 2016 Ayke van Laethem.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package gocount

import (
	"runtime"
	"testing"
)

func TestNumber(t *testing.T) {
	const defaultNumber = 2 // number by default in the testing package, subject to change

	if n := Number(); n != defaultNumber {
		t.Errorf("wrong number when starting test, expected %d, got %d", defaultNumber, n)
	}

	block := make(chan struct{})      // blocking goroutine, until the end of the test
	countStart := make(chan struct{}) // make sure goroutines are started
	countStop := make(chan struct{})  // make sure goroutines are stopped

	const extraNumber = 5

	for i := 0; i < extraNumber; i++ {
		go func() {
			countStart <- struct{}{}
			<-block
			countStop <- struct{}{}
		}()
		<-countStart // make sure the goroutine is started

		expected := defaultNumber + i + 1
		actual := Number()
		if actual != expected {
			t.Errorf("wrong number after starting goroutine #%d, expected %d, got %d", i+1, expected, actual)
		}
	}

	// stop all goroutines
	close(block)
	for i := 0; i < extraNumber; i++ {
		// make sure the goroutine is stopped
		<-countStop
		runtime.Gosched()
	}

	if n := Number(); n != defaultNumber {
		t.Errorf("wrong number when ending test, expected %d, got %d", defaultNumber, n)
	}
}
