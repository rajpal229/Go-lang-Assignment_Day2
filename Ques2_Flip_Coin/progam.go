package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Ques2
	var n int
	fmt.Print("Enter Number of times to run the Stimulation: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		rn := rand.Intn(2-0) + 0
		switch rn {
		case 0:
			fmt.Println("Head")
		case 1:
			fmt.Println("Tail")
		}
	}
}
