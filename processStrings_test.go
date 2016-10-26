package main

import "testing"

func TestProcessStrings(t *testing.T) {
	ss := []string{"a;a,b",
		"",
		"aa,b;cc",
	}

	expectedSS := []float64{
		0,
		-1,
		4.5,
	}

	rv := processStrings(ss)

	if rv[0] != expectedSS[0] {
		t.Fatalf("Max total SS score of \"%v\" is \"%v\"", ss[0], expectedSS[0])
	}

	if rv[1] != expectedSS[1] {
		t.Fatalf("Max total SS score of \"%v\" is \"%v\"", ss[1], expectedSS[1])
	}

	if rv[2] != expectedSS[2] {
		t.Fatalf("Max total SS score of \"%v\" is \"%v\"", ss[2], expectedSS[2])
	}
}
