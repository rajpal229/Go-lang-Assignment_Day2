package main

import (
	"fmt"
)

func main() {
	//Ques9
	var a, b, c int
	fmt.Println("Enter sides of the Triangle: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a == b && a == c {
		fmt.Println("Equilateral")
	} else if a == b || b == c || c == a {
		fmt.Println("Isoceles")
	} else {
		fmt.Println("Scalene")
	}
}
