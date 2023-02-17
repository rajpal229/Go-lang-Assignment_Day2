package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Ques6
	var n int
	fmt.Print("Enter the Number: ")
	fmt.Scan(&n)
	var op_string string
	if n%3 == 0 {
		op_string = op_string + "Pling"
	}
	if n%5 == 0 {
		op_string = op_string + "Plang"
	}
	if n%7 == 0 {
		op_string = op_string + "Plong"
	}
	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		op_string = strconv.Itoa(n)
	}
	fmt.Println(op_string)
}
