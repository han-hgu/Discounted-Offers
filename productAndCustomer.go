package main

import "fmt"

type namedObj struct {
	name string
	cs   *charSet
}

type product struct {
	*namedObj
}

type customer struct {
	*namedObj
	numVowel int
}

func newObj(name string) (*namedObj, error) {
	cs, err := set(name)

	if err != nil {
		return nil, err
	}

	nobj := new(namedObj)
	nobj.name = name
	nobj.cs = cs

	return nobj, nil
}

//NewCustomer creates a new customer, this method should
//be called if creating a new customer
func NewCustomer(name string) (*customer, error) {
	nobj, err := newObj(name)
	if err != nil {
		return nil, err
	}

	nc := new(customer)
	nc.numVowel = NumOfVowels(name)
	nc.namedObj = nobj

	return nc, nil
}

func NewProduct(name string) (*product, error) {
	nobj, err := newObj(name)
	if err != nil {
		return nil, err
	}

	nc := new(product)
	nc.namedObj = nobj
	return nc, nil
}

//Assume the comparison is case insensitive
func calculateSS(p *product, c *customer) float64 {
	//If either name is empty, return 0, it is a dummy product or customer
	if p.name == "" || c.name == "" {
		return 0
	}

	var rv float64
	if len(p.name)%2 == 0 {
		rv = float64(c.numVowel) * vowelFactor
	} else {
		rv = float64((len(c.name) - c.numVowel)) * consonantFactor
	}

	return rv
}

//TODO: check capitalize
func Intersects(p *product, c *customer) bool {
	return p.cs.intersects(c.cs)
}

func fillDummyCustomers(cs *[]*customer, num int) {
	if num <= 0 {
		return
	}

	for i := 0; i < num; i++ {
		nc, _ := NewCustomer("")
		*cs = append(*cs, nc)
	}
}

func fillDummyProducts(ps *[]*product, num int) {
	fmt.Println("HAN >>>>> fillDummyProducts called. ", num)
	if num <= 0 {
		return
	}

	for i := 0; i < num; i++ {
		np, _ := NewProduct("")
		*ps = append(*ps, np)
	}
}
