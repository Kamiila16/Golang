package main

import "fmt"

//ex1,2,3
func main() {
	fmt.Println("Hello World!")
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Println("positive")
	} else if n == 0 {
		fmt.Println("zero")
	} else if n < 0 {
		fmt.Println("negative")
	}
	var sumn int = 0
	for i := 0; i < 11; i++ {
		sumn = sumn + i
	}
	fmt.Println(sumn)
	var d int
	fmt.Scan(&d)
	switch d {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	}

}
