package main

import (
	"fmt"
	"testing"
)

func TestNewDiscountOfferMatrix(t *testing.T) {
	//TODO: This one has user HAN GU in it!!
	//TODO: documentation: leading or trailing space will be part of the customer or product name
	l := "Jack Abraham,John Evans,Ted Dziuba,Han Gu;iPad 2 - 4-pack,Girl Scouts Thin Mints,Nerf Crossbow"
	dsm, _ := NewDiscountOfferMatrix(l)
	fmt.Println("HAN >>> row", dsm.customers)

	for _, v := range dsm.customers {
		fmt.Println(v.name)
	}
	fmt.Println("HAN >>> column", dsm.products)

	for _, v := range dsm.products {
		fmt.Println(v.name)
	}
}
