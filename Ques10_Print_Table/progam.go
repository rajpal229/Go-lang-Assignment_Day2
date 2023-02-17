package main

import (
	"fmt"
)

func main() {
	//Ques10
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	for {
		if n >= 2 || n <= 25 {
			for i := 0; i < 10; i++ {
				fmt.Printf("%v x %v = %v\n", n, i+1, n*(i+1))
			}
			break
		} else {
			fmt.Println("Enter value between 2 and 25")
		}
	}
}
