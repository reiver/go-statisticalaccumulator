# Go Statistical Accumulator

A Go library for efficiently calculating the statistical *mean* and *variance* of a data set, **in a single pass**.

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

	// Add some integers, as strings.
	statisticalAccumulator.PutString("17")
	statisticalAccumulator.PutString("19")
	statisticalAccumulator.PutString("18")
	statisticalAccumulator.PutString("17")
	statisticalAccumulator.PutString("19")
	statisticalAccumulator.PutString("-12")
	statisticalAccumulator.PutString("0")

	// Add some numbers, as string, that aren't integers, but instead have decimal points
	statisticalAccumulator.PutString("3.14159265358979323846264338327950")
	statisticalAccumulator.PutString("18725.2200440088")
	statisticalAccumulator.PutString("-321.0123")

	// Add some numbers, as string, that aren't integers, but are fractions.
	statisticalAccumulator.PutString("1/3")
	statisticalAccumulator.PutString("22/7")
	statisticalAccumulator.PutString("-7/11")

	// Add some numbers that are in (an efficient) binary (format as an) int64.
	statisticalAccumulator.PutInt64(1234567)
	statisticalAccumulator.PutInt64(0)
	statisticalAccumulator.PutInt64(-2)


	n        := statisticalAccumulator.N()
	mean     := statisticalAccumulator.Mean()
	variance := statisticalAccumulator.Variance()


	// Given the PUT calls above, prints: "The amount of data is N = 5/1".
	//
	// Note that statisticalAccumulator.N() returns a *big.Rat,
	// thus why the outputted value is "5/1" and not just "5".
	// (By default, big.Rat string serialization is as a fraction.)
	fmt.Printf("The amount of data is N = %v\n", n)

	// Given the PUT calls above, prints: "The mean is 18/1".
	//
	// Note that statisticalAccumulator.Mean() returns a *big.Rat,
	// thus why the outputted value is "18/1" and not just "18".
	// (By default, big.Rat string serialization is as a fraction.)
	fmt.Printf("The mean is %v\n", mean)

	// Given the PUT calls above, prints: "The variance is 4/5".
	fmt.Printf("The variance is %v\n", variance)
}
```
