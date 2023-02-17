package main

import (
	"fmt"
)

func main() {
	//Ques1
	var inp int
	fmt.Println("Welcome to Temperature Unit Converter Program.")
	fmt.Println("Enter\n0. Centigrade to Fahrenheit\n1. Fahrenheit to Centigrade")
	fmt.Scan(&inp)
	switch inp {
	case 0:
		var c float32
		fmt.Print("Enter Centigrade: ")
		fmt.Scan(&c)
		f := 1.8*c + 32
		fmt.Printf("%v degree Centigrade is %v degree Fahrenheit", c, f)
	case 1:
		var f float32
		fmt.Print("Enter Fahrenheit: ")
		fmt.Scan(&f)
		c := (f - 32) / 1.8
		fmt.Printf("%v degree Fahrenheit is %v degree Centigrade", f, c)
	}
}
