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


// The fmt.Scan() function is a great function for easily fetching & using user input through the command line.

// But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.

// For the moment, we only need single words / digits as input, so that's no problem.

// Later in the course, when we work on projects where more complex input values are required, you'll therefore learn about an alternative to fmt.Scan().