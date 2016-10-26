package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/clyphub/munkres"
)

// discountOfferMatrix encapsulates the munkres package with
// customers and products info
type discountOfferMatrix struct {
	customers []*customer
	products  []*product
	m         *munkres.Matrix
}

// NewDiscountOfferMatrix translates the input string into product/customer
// information; It also stores the SS scores for product/customer combination
// Assumption: the first ';' will be used as the delimiter, if more than one
// is in the string, they become part of the product information
func NewDiscountOfferMatrix(line string) *discountOfferMatrix {
	rawl := strings.SplitN(line, ";", 2)
	var rawCustomers, rawProducts string
	if len(rawl) < 2 {
		rawCustomers = rawl[0]
		rawProducts = ""
	} else {
		rawCustomers = rawl[0]
		rawProducts = rawl[1]
	}

	dsm := new(discountOfferMatrix)

	for _, rc := range strings.Split(rawCustomers, ",") {
		dsm.customers = append(dsm.customers, NewCustomer(rc))
	}

	for _, rp := range strings.Split(rawProducts, ",") {
		dsm.products = append(dsm.products, NewProduct(rp))
	}

	// Add filler to make the matrix "square"
	if len(dsm.products) > len(dsm.customers) {
		fillDummyCustomers(&dsm.customers, len(dsm.products)-len(dsm.customers))
	} else if len(dsm.products) < len(dsm.customers) {
		fillDummyProducts(&dsm.products, len(dsm.customers)-len(dsm.products))
	}

	dsm.m = constructMatrixHelper(dsm.customers, dsm.products)

	return dsm
}

// computeMaxTotalSS computes the maximum total score based on the info stored
// in dsm
func (dsm *discountOfferMatrix) computeMaxTotalSS() float64 {
	var totalScore int64 = 0

	// calculate the max using the Hungarian Method
	rc := munkres.ComputeMunkresMax(dsm.m)

	for _, v := range rc {
		// munkres.RowCol doesn't export its row and col members hence
		// the below workaround. Ideally the library should provide us a way
		// to retrieve the min/max value
		s := fmt.Sprintln(v)
		re := regexp.MustCompile("[0-9]+")
		rowCol := re.FindAllString(s, -1)
		if len(rowCol) != 2 {
			// Don't bother returning as it is unlikely any other worker could
			// produce useful result
			log.Fatal("munkres package returns an uninterpretable struct")
		}

		row, err := strconv.Atoi(rowCol[0])
		if err != nil {
			log.Fatal("munkres package returns an uninterpretable struct")
		}

		col, err := strconv.Atoi(rowCol[1])
		if err != nil {
			log.Fatal("munkres package returns an uninterpretable struct")
		}

		index := row*len(dsm.customers) + col
		totalScore += dsm.m.A[index]
	}

	return float64(totalScore) / expandFactor
}

// constructMatrixHelper constructs the munkres.Matrix based on the customer and product info
func constructMatrixHelper(cs []*customer, ps []*product) *munkres.Matrix {
	if len(cs) != len(ps) {
		// This is unlikely to occur hence generating fatal error.
		// This method should only be used as a private helper function.
		// We explicitly make the two slices of equal length.
		log.Fatal("Unable to construct a munkres matrix with unequal number of customers and products")
	}

	mm := munkres.NewMatrix(len(cs))

	i := 0
	for _, c := range cs {
		for _, p := range ps {
			ss := calculateSS(p, c)
			mm.A[i] = ss
			i += 1
		}
	}

	return mm
}
