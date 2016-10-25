package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/clyphub/munkres"
	"github.com/gammazero/workerpool"
)

var result []float64

//TODO: Caching necessary?
/*
var productCache map[string]*product
var customerCache map[string]*customer
*/

const vowelFactor = 1.5
const consonantFactor = 1

//test no customer is provided in the list
func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var rs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rs = append(rs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rs, nil
}

type discountOfferMatrix struct {
	customers []*customer
	products  []*product
	*munkres.Matrix
}

//compute returns -1 if error is encountered processing this line
func compute(line string) float64 {
	//_ := NewDiscountOfferMatrix(line)
	//compute()
	return -1
}

//transform transforms the line to the customer and products
func NewDiscountOfferMatrix(line string) (*discountOfferMatrix, error) {
	rawl := strings.Split(line, ";")
	rawCustomers := rawl[0]
	rawProducts := rawl[1]

	dsm := new(discountOfferMatrix)

	for _, rc := range strings.Split(rawCustomers, ",") {
		nc, err := NewCustomer(rc)
		if err != nil {
			return nil, err
		}
		dsm.customers = append(dsm.customers, nc)
	}

	for _, rp := range strings.Split(rawProducts, ",") {
		np, err := NewProduct(rp)
		if err != nil {
			return nil, err
		}
		dsm.products = append(dsm.products, np)
	}

	//TODO: test one customer and one product
	//Add filler, calculateSS knows how to calculate such
	if len(dsm.products) > len(dsm.customers) {
		fillDummyCustomers(dsm.customers, len(dsm.products)-len(dsm.customers))
	} else if len(dsm.products) < len(dsm.customers) {
		fillDummyProducts(&dsm.products, len(dsm.customers)-len(dsm.products))
	}

	var err error
	dsm.Matrix, err = constructMatrixHelper(dsm.customers, dsm.products)
	if err != nil {
		return nil, err
	}

	return dsm, nil
}

func constructMatrixHelper(cs []*customer, ps []*product) (*munkres.Matrix, error) {
	if len(cs) != len(ps) {
		return nil, errors.New("Unable to construct a munkres matrix with unequal number of customers and products")
	}

	mm := munkres.NewMatrix(len(cs))

	return mm, nil
}

//parse the input file and create worker for each line
func process(info []string, worker_num int) []float64 {
	fmt.Println("worker number:", worker_num)

	rv := make([]float64, len(info))
	//for each line, send it to a worker
	wp := workerpool.New(worker_num)

	for i := 0; i < len(info); i++ {
		wp.Submit(func() {
			//TODO: Test empty line in the input file
			if info[i] != "" {
				r := compute(info[i])
				rv[i] = r
			}
		})
	}

	wp.Stop()

	return rv
}
