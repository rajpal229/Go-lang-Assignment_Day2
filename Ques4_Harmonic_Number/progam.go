package main

import (
	"fmt"
)

func main() {
	//Ques4
	var n int
	var hn float64
	fmt.Print("Enter Number: ")
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		hn = hn + 1/float64(i)
	}
	fmt.Printf("%.5f", hn)
}
