package main

import (
	"fmt"
	"math"
)

func main() {
	//Ques8
	var op int
	var (
		p           float64 = 1000
		rr1         float64 = 6.25
		rr2         float64 = 7.50
		rr3         float64 = 9.5
		t           float64 = 2.5
		finalAmount float64
	)
	fmt.Println("Select interest rate")
	fmt.Println(" 1.", rr1, "%\n", "2.", rr2, "%\n", "3.", rr3, "%")
	switch fmt.Scan(&op); op {
	case 1:
		finalAmount = p * (math.Pow(1+(rr1/1), (1 * t)))
	case 2:
		finalAmount = p * (math.Pow(1+(rr2/1), (1 * t)))
	case 3:
		finalAmount = p * (math.Pow(1+(rr3/1), (1 * t)))
	}
	fmt.Print("Final Amount is: Rs ")
	fmt.Printf("%.2f", finalAmount)
}
