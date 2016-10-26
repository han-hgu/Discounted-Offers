package main

import "testing"

func TestNewCustomer(t *testing.T) {
	cname1 := "John Smith"
	nc := NewCustomer(cname1)

	if nc.name != cname1 {
		t.Fatal("Customer name set correctly")
	}

	if nc.numVowel != NumOfVowels(cname1) {
		t.Fatal("Customer numOfVowel set correctly")
	}
}

func TestCalculateSS_EmptyCustomerName_EmptyProductName(t *testing.T) {
	nce := NewCustomer("")
	npe := NewProduct("")
	nc := NewCustomer("customer")
	np := NewProduct("product")

	if calculateSS(np, nce) != 0 {
		t.Fatal("If customer name is empty, the SS score is 0")
	}

	if calculateSS(npe, nc) != 0 {
		t.Fatal("If product name is empty, the SS score is 0")
	}

	if calculateSS(npe, nce) != 0 {
		t.Fatal("If customer and product name are empty, the SS score is 0")
	}
}

func TestCalculateSS_rule1(t *testing.T) {
	nc := NewCustomer("customer")
	np := NewProduct("product-a")
	expectedSS := 3 * vowelFactor * gcdFactor * expandFactor

	if float64(calculateSS(np, nc)) != expectedSS {
		t.Fatalf("Customer \"%v\" and product \"%v\" has SS score of \"%v\"", nc.name, np.name, expectedSS)
	}
}

func TestCalculateSS_rule2(t *testing.T) {
	nc := NewCustomer("customer-aa")
	np := NewProduct("product-ab")
	expectedSS := float64(5 * consonantFactor * expandFactor)

	if float64(calculateSS(np, nc)) != expectedSS {
		t.Fatalf("Customer \"%v\" and product \"%v\" has SS score of \"%v\"", nc.name, np.name, expectedSS)
	}
}

func TestCalculateSS_rule3(t *testing.T) {
	nc := NewCustomer("customer-a")
	np := NewProduct("product-ab")
	expectedSS := float64(5 * consonantFactor * gcdFactor * expandFactor)

	if float64(calculateSS(np, nc)) != expectedSS {
		t.Fatalf("Customer \"%v\" and product \"%v\" has SS score of \"%v\"", nc.name, np.name, expectedSS)
	}
}

func TestFillDummyCustomers(t *testing.T) {
	var cs []*customer
	cs = append(cs, NewCustomer("abc"))

	dn := 2
	fillDummyCustomers(&cs, dn)
	if len(cs) != 3 {
		t.Fatal("Number of customers is correct after filling the dummy customers")
	}

	if cs[0].name != "abc" {
		t.Fatal("The original customer is not modified")
	}

	if cs[1].name != "" {
		t.Fatal("The first dummy customer is added with empty name")
	}

	if cs[2].name != "" {
		t.Fatal("The second dummy customer is added with empty name")
	}
}

func TestFillDummyProducts(t *testing.T) {
	var cs []*product
	cs = append(cs, NewProduct("abc"))

	dn := 2
	fillDummyProducts(&cs, dn)
	if len(cs) != 3 {
		t.Fatal("Number of products is correct after filling the dummy products")
	}

	if cs[0].name != "abc" {
		t.Fatal("The original product is not modified")
	}

	if cs[1].name != "" {
		t.Fatal("The first dummy product is added with empty name")
	}

	if cs[2].name != "" {
		t.Fatal("The second dummy product is added with empty name")
	}
}
