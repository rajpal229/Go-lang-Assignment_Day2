package main

import (
	"fmt"
	"strings"
)

func main() {
	//Ques7 ??using and n or instead nested loop
	var a, e, p string
	fmt.Print("Enter Attendance-\nEnter P - For Present\nEnter A - For Absent\na is: ")
	fmt.Scan(&a)
	fmt.Print("e is: ")
	fmt.Scan(&e)
	fmt.Print("p is: ")
	fmt.Scan(&p)
	if strings.ToUpper(a) == "P" && strings.ToUpper(e) == "P" && strings.ToUpper(p) == "P" {
		fmt.Println("All Present")
	} else if strings.ToUpper(a) == "P" || strings.ToUpper(e) == "P" || strings.ToUpper(p) == "P" {
		fmt.Println("One or more - Present")
	} else {
		fmt.Println("None Present")
	}
}
