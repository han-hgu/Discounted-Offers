package main

import "testing"

func TestNewDiscountOfferMatrix_emptyString(t *testing.T) {
	dsm := NewDiscountOfferMatrix("")

	if dsm == nil {
		t.Fatal("Discount offer matrix could be created with empty string input")
	}

	if len(dsm.customers) != 1 {
		t.Fatal("Customer list is empty if input string is empty")
	}

	if len(dsm.products) != 1 {
		t.Fatal("Product list is empty if input string is empty")
	}

	if len(dsm.m.A) != 1 {
		t.Fatal("Munkres matrix of size 1 is created")
	}
}

func TestNewDiscountOfferMatrix_OnlyCustomers(t *testing.T) {
	line := "John Smith, Mike Two, Tylor Wong;"
	dsm := NewDiscountOfferMatrix(line)

	if dsm == nil {
		t.Fatal("Discount offer matrix could be created with customer only string input")
	}

	if len(dsm.customers) != 3 {
		t.Fatal("Customer list has the correct size")
	}

	if len(dsm.products) != 3 {
		t.Fatal("Product list is empty")
	}

	if len(dsm.m.A) != 9 {
		t.Fatal("Munkres matrix is constructed if input string contains only customers")
	}
}

func TestNewDiscountOfferMatrix_OnlyProducts(t *testing.T) {
	line := ";Ipad - 40 packs, Plasma TV"
	dsm := NewDiscountOfferMatrix(line)

	if dsm == nil {
		t.Fatal("Discount offer matrix could be created with customer only string input")
	}

	if len(dsm.customers) != 2 {
		t.Fatal("Customer list has the correct size")
	}

	if len(dsm.products) != 2 {
		t.Fatal("Product list is empty")
	}

	if len(dsm.m.A) != 4 {
		t.Fatal("Munkres matrix is constructed if input string contains only products")
	}
}

func TestNewDiscountOfferMatrix_ProductsAndCustomers(t *testing.T) {
	line := "John Smith, Mike Two, Tylor Wong; Ipad - 40 packs, Plasma TV"
	dsm := NewDiscountOfferMatrix(line)

	if dsm == nil {
		t.Fatal("Discount offer matrix could be created with customer only string input")
	}

	if len(dsm.customers) != 3 {
		t.Fatal("Customer list has the correct size")
	}

	if len(dsm.products) != 3 {
		t.Fatal("Product list has the correct size")
	}

	if len(dsm.m.A) != 9 {
		t.Fatal("Munkres matrix is constructed if input string contains only products")
	}
}

func TestComputeMaxTotalSS(t *testing.T) {
	line := "a; bc"
	expectedSS := 1.5
	dsm := NewDiscountOfferMatrix(line)
	if dsm.computeMaxTotalSS() != expectedSS {
		t.Fatalf("Max total SS score of \"%v\" is \"%v\"", line, expectedSS)
	}

	line = "ab;bc"
	dsm = NewDiscountOfferMatrix(line)
	expectedSS = 2.25
	if dsm.computeMaxTotalSS() != expectedSS {
		t.Fatalf("Max total SS score of \"%v\" is \"%v\"", line, expectedSS)
	}

	line = "aa; bc"
	dsm = NewDiscountOfferMatrix(line)
	expectedSS = 4.5
	if dsm.computeMaxTotalSS() != expectedSS {
		t.Fatalf("Max total SS score of \"%v\" is \"%v\"", line, expectedSS)
	}
}

func TestConstructMatrixHelper(t *testing.T) {
	var cs []*customer
	var ps []*product

	nc1 := NewCustomer("aeiou")
	nc2 := NewCustomer("bcdfp")
	cs = append(cs, nc1, nc2)
	np1 := NewProduct("p1")
	np2 := NewProduct("P2")
	ps = append(ps, np1, np2)
	m := constructMatrixHelper(cs, ps)

	if len(m.A) != 4 {
		t.Fatalf("The length of the matrix with 2 customers and 2 products is 4")
	}
}
