package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	 expectedReturnRate := 5.5
	var years float64 = 10

	const inflation float64= 5.5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}
