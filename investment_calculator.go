package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5 // if we will not give any value to exprectedReturnRate then it will take 5.5 by default.
	var years float64 = 10
	const inflationRate = 2.5

	fmt.Print("investment Amout: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	//  if you are not assigning a value to a variable then you have to declare variable with var keyward and should have to give its type .
	//  you can't declare a variable with := if you are not assigning the value to a variable.

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
