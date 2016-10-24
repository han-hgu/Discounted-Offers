package main

import (
	"errors"
	"strings"
	"unicode"
)

//TODO: better name?
type charset struct {
	bits uint32
}

func (c *charSet) set(i byte) {
	c.bits |= 1 << (i - 'a')
}

func (c1 *charSet) intersects(c2 *charSet) bool {
	return c1.bits&c2.bits != 0
}

//Compared case insensitive
func set(s string) (error, charSet) {
	ls := strings.ToLower(s)
	var rcs charSet
	for _, c := range ls {
		if c > unicode.MaxASCII {
			//TODO: Make error display c
			return errors.New("Only ASCII string can not processed")
		}
		rcs.set(ls[i])
	}

	return nil, rcs
}
