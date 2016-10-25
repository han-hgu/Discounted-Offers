package main

func NumOfVowels(s string) int {
	rv := 0
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, c := range s {
		for _, v := range vowels {
			if v == c {
				rv += 1
			}
		}
	}

	return rv
}
