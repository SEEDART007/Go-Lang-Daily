package main

import (
	"fmt"
	"math"
)

func main() {
	var investment float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investment * math.Pow(1+expectedReturnRate/100, years)
	fmt.Print(futureValue)
}
