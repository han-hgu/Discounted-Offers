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

func TestIntersects(t *testing.T) {
	var cs1, cs2, cs3, cs4 charSet

	cs1.set('a')
	cs2.set('a')
	if !cs1.intersects(&cs2) {
		t.Fatal("Identical charSets intersect each other")
	}

	cs3.set('b')
	cs3.set('b')
	cs4.set('z')
	if cs3.intersects(&cs4) {
		t.Fatal("Different charSets doesn't intersect")
	}
}
