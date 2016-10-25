package main

import (
	"strings"
	"unicode"
)

//TODO: better name?
type charSet struct {
	bits uint32
}

func (c *charSet) set(i byte) {
	c.bits |= 1 << (i - 'a')
}

func (c1 *charSet) intersects(c2 *charSet) bool {
	return c1.bits&c2.bits != 0
}

//Compared case insensitive
func set(s string) (*charSet, error) {
	ls := strings.ToLower(s)
	rcs := new(charSet)
	for _, c := range ls {
		if c > unicode.MaxASCII {
			//TODO: Make error display c
			return nil, nil
		}
		rcs.set(byte(c))
	}

	return rcs, nil
}
