package main

import (
	"fmt"
	"math"
)

func main() {
	//Ques5
	var p, y, r float64
	fmt.Println("Enter Vales:")
	fmt.Print("Principal(P): ")
	fmt.Scan(&p)
	fmt.Print("Number of Years(Y): ")
	fmt.Scan(&y)
	fmt.Print("Rate of Interest(R): ")
	fmt.Scan(&r)
	payment := (p * r / (12 * 100)) / (1 - math.Pow((1+r/(12*100)), (-12*5)))
	fmt.Print("EMI to be paid: ")
	fmt.Printf("%.2f", payment)
}
