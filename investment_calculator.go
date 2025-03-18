package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64 = 10
	const inflationRate = 2.5
	fmt.Print("investment Amout: ")
	fmt.Scan(&investmentAmount)

	// const values can't be changes later and it should be initialized at the time of decalration.
	// variables declared with var keyward can be decalred first and initilized later.
	

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
