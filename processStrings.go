package main

import "sync"

// processStrings treats each entry in the info slice as a product/customer set
// and calculates the maximum total SS score for that entry; it returns a slice
// with each entry storing the maximum total SS score corresponding to the input
// entry
func processStrings(info []string) []float64 {
	rv := make([]float64, len(info))
	var wg sync.WaitGroup

	// create a worker for every line
	for i := 0; i < len(info); i++ {
		if info[i] != "" {
			wg.Add(1)

			//no need to lock since every thread is modifying
			//a different slot of the array
			go func(line string, index int, rs *[]float64) {
				defer wg.Done()
				r := NewDiscountOfferMatrix(line).computeMaxTotalSS()
				(*rs)[index] = r
			}(info[i], i, &rv)
		} else {
			rv[i] = -1
		}
	}

	wg.Wait()
	return rv
}
