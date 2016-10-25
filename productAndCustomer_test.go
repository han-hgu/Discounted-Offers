package main

import "testing"

func TestNewCustomer(t *testing.T) {
	cname1 := "John Smith"
	nc, err := NewCustomer(cname1)

	if err != nil {
		t.Fatal("Customer instance created successfully")
	}

	if nc.name != cname1 {
		t.Fatal("Customer name set correctly")
	}

	if nc.numVowel != NumOfVowels(cname1) {
		t.Fatal("Customer numOfVowel set correctly")
	}

	cs1, _ := set(cname1)
	if nc.cs.bits != cs1.bits {
		t.Fatal("Customer charSet set correctly")
	}
}

func TestNewProduct(t *testing.T) {

}
