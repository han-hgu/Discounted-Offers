package main

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	var cs charSet
	cs.set('a')
	if cs != 'a' {
		fmt.Println("HAN >>>>>", cs)
	}
}
