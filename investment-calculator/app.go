package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var (
		investAmount    float64
		years           float64
		interest        float64
		futureValue     float64
		futureRealValue float64
	)
	fmt.Print("Enter the investment Amount: ")
	fmt.Scan(&investAmount)
	fmt.Print("\nEnter number of Years: ")
	fmt.Scan(&years)
	fmt.Print("\nEnter the interest rate: ")
	fmt.Scan(&interest)

	futureValue = investAmount * math.Pow(1+interest/100, years)
	futureRealValue = futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println("\nThe future value will be: ", futureValue)
	fmt.Println("\nThe future real value will be: ", futureRealValue)

}
