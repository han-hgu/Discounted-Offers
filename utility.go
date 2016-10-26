package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ParseFile parses a file and returns a slice of strings
func ParseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var rs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rs = append(rs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rs, nil
}

// NumOfVowels calculates the number of vowels in s
func NumOfVowels(s string) int {
	rv := 0
	sl := strings.ToLower(s)
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	for _, c := range sl {
		for _, v := range vowels {
			if v == c {
				rv += 1
			}
		}
	}

	return rv
}

// NumOfLetters calculates the number of letters in s
func NumOfLetters(s string) int {
	rv := 0

	for _, c := range s {
		if unicode.IsLetter(c) {
			rv += 1
		}
	}

	return rv
}

// gcd calculates the greatest common divisor between x and y
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// print prints each entry to two decimal places, it skips all
// the entries with value -1
func print(d []float64) {
	for _, v := range d {
		if v != -1 {
			fmt.Printf("%.2f\n", v)
		}
	}
}
