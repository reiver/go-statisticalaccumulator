# Go Statistical Accumulator

A Go library for efficiently calculating the statistical *mean* and *variance* of the data, **in a single pass**.

## Usage
Usage of this library is like in the following example:

```
package main

import (
	"fmt"
	"github.com/reiver/go-statisticalaccumulator"
)

func main() {

	statisticalAccumulator := statisticalaccumulator.New()

	statisticalAccumulator.PutString("17")
	statisticalAccumulator.PutString("19")
	statisticalAccumulator.PutString("18")
	statisticalAccumulator.PutString("17")
	statisticalAccumulator.PutString("19")


	// Given the PUT calls above, prints: "The amount of data is N = 5/1".
	//
	// Note that statisticalAccumulator.N() returns a *big.Rat,
	// thus why the outputted value is "5/1" and not just "5".
	// (By default, big.Rat string serialization is as a fraction.)
	fmt.Printf("The amount of data is N = %v\n", statisticalAccumulator.N())

	// Given the PUT calls above, prints: "The mean is 18/1".
	//
	// Note that statisticalAccumulator.Mean() returns a *big.Rat,
	// thus why the outputted value is "18/1" and not just "18".
	// (By default, big.Rat string serialization is as a fraction.)
	fmt.Printf("The mean is %v\n", statisticalAccumulator.Mean())

	// Given the PUT calls above, prints: "The variance is 4/5".
	fmt.Printf("The variance is %v\n", statisticalAccumulator.Variance())
}
```
