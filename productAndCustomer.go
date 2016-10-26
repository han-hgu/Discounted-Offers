package main

const vowelFactor = 1.5
const consonantFactor = 1
const gcdFactor = 1.5
const expandFactor = 1000

type namedObj struct {
	name         string
	numOfLetters int
}

type product struct {
	*namedObj
}

type customer struct {
	*namedObj
	numVowel int
}

// newobj creates a namedObj instance
func newObj(name string) *namedObj {
	nobj := new(namedObj)
	nobj.name = name
	nobj.numOfLetters = NumOfLetters(name)

	return nobj
}

// NewCustomer creates a new customer
func NewCustomer(name string) *customer {
	nobj := newObj(name)
	nc := new(customer)
	nc.numVowel = NumOfVowels(name)
	nc.namedObj = nobj
	return nc
}

// NewProduct creates a new product
func NewProduct(name string) *product {
	nobj := newObj(name)
	np := new(product)
	np.namedObj = nobj
	return np
}

// calculateSS calculates the SS score for product and customer provided,
// assuming the ascii character comparison is case insensitive.
func calculateSS(p *product, c *customer) int64 {
	//If either name is empty, return 0 for dummy product/customer
	if p.name == "" || c.name == "" {
		return 0
	}

	var rv float64
	if p.numOfLetters%2 == 0 { //rule #1
		rv = float64(c.numVowel) * vowelFactor
	} else { //rule #2
		rv = float64(c.numOfLetters-c.numVowel) * consonantFactor
	}

	if gcd(c.numOfLetters, p.numOfLetters) > 1 { //rule #3
		rv = rv * gcdFactor
	}

	// munkres package can only calculate matrices with entry type int,
	// expand the float 1000 times and take the int portion to achieve
	// the accuracy level of "2 digits after decimal point"
	return int64(rv * expandFactor)
}

// fillDummyCustomers adds the specified number of nameless customers
// into the customer slice, the SS score will always be 0 if the dummy
// customer is involved
func fillDummyCustomers(cs *[]*customer, num int) {
	for i := 0; i < num; i++ {
		*cs = append(*cs, NewCustomer(""))
	}
}

// fillDummyProducts adds the specified number of nameless products
// into the product slice
func fillDummyProducts(ps *[]*product, num int) {
	for i := 0; i < num; i++ {
		*ps = append(*ps, NewProduct(""))
	}
}
