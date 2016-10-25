package main

import "testing"

func TestNumOfVowels(t *testing.T) {
	s1 := "abcdefg"

	if NumOfVowels(s1) != 2 {
		t.Fatalf("Number of vowerls in %v is 2", s1)
	}
}
