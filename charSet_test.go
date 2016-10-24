package main

import (
	"math"
	"testing"
)

func TestSet(t *testing.T) {
	var cs1, cs2 charSet

	cs1.set('a')
	if cs1.bits != 1 {
		t.Fatalf("charSet has the 0th bit set for 'a'")
	}

	cs2.set('z')
	if float64(cs2.bits) != math.Pow(2, 'z'-'a') {
		t.Fatalf("charSet has the 25th bit set for 'z'")
	}
}
