package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNumOfVowels(t *testing.T) {
	s1 := "Annie Maggie TAU -- "
	expectedNumV := 8

	if NumOfVowels(s1) != expectedNumV {
		t.Fatalf("Number of vowerls in \"%v\" is \"%v\"", s1, expectedNumV)
	}
}

func TestNumOfLetters(t *testing.T) {
	s1 := ""
	expectedNumL := 0

	if NumOfLetters(s1) != expectedNumL {
		t.Fatalf("Number of letters in \"%v\" is \"%v\"", s1, expectedNumL)
	}

	s1 = "ab ewew& - ;wew"
	expectedNumL = 9

	if NumOfLetters(s1) != expectedNumL {
		t.Fatalf("Number of letters in \"%v\" is \"%v\"", s1, expectedNumL)
	}
}

func Test_gcd(t *testing.T) {
	if gcd(1, 5) != 1 {
		t.Fatalf("GCD of 1 and 5 is 1")
	}

	if gcd(4, 8) != 4 {
		t.Fatalf("GCD of 4 and 8 is 4")
	}
}

func TestParseFile(t *testing.T) {
	if _, err := ParseFile(""); err == nil {
		t.Fatal("ParseFile() returns error if file name doesn't exist")
	}

	file, _ := ioutil.TempFile(os.TempDir(), "test")
	defer os.Remove(file.Name())

	writeString := []string{
		"write1\n",
		"write2\n",
		"write3\n",
	}

	expectedString := []string{
		"write1",
		"write2",
		"write3",
	}
	f, _ := os.Create(file.Name())
	defer f.Close()
	f.WriteString(writeString[0])
	f.WriteString(writeString[1])
	f.WriteString(writeString[2])

	rv, _ := ParseFile(file.Name())
	if len(rv) != len(expectedString) {
		t.Fatal("The number of elements in the parsed result is equal to the number of lines in the input file")
	}

	if rv[0] != expectedString[0] {
		t.Fatal("The first element in the parsed result is correct")
	}

	if rv[1] != expectedString[1] {
		t.Fatal("The first element in the parsed result is correct")
	}

	if rv[2] != expectedString[2] {
		t.Fatal("The first element in the parsed result is correct")
	}
}
