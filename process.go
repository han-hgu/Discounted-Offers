package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/clyphub/munkres"
	"github.com/gammazero/workerpool"
)

var result []float32

const vowelFactor float32 = 1.5
const consonantFactor float32 = 1

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

type customer struct {
	name     string
	numVowel int
	cs       *charSet
}

type product struct {
	name string
	cs   *charSet
}

//Assume the comparison is case insensitive
func calculateSS(p *product, c *customer) float32 {
	var rv float32
	if len(p.name)%2 == 0 {
		rv = float32(c.numVowel) * vowelFactor
	} else {
		rv = float32((len(c.name) - c.numVowel)) * consonantFactor
	}

	return rv
}

func Intersects(p *product, c *customer) bool {
	return p.cs.intersects(c.cs)
}

type DiscountOfferMatrix struct {
	rows    []customer
	columns []product
	munkres.Matrix
}

func compute(line string) float32 {
	transform(line)
	//compute()
	return 0
}

func transform(input string) *munkres.Matrix {
	fmt.Println(input)
	return nil
}

//parse the input file and create worker for each line
func process(info []string, worker_num int) []float32 {
	fmt.Println("worker number:", worker_num)

	rv := make([]float32, len(info))
	//for each line, send it to a worker
	wp := workerpool.New(worker_num)

	for i := 0; i < len(info); i++ {
		wp.Submit(func() {
			r := compute(info[i])
			rv[i] = r
		})
	}

	wp.Stop()

	return rv
}
