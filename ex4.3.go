package main

import "fmt"

func qr(a int, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
func main() {
	var n1, n2 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	q, r := qr(n1, n2)
	fmt.Println("quotient = ", q, "remainder = ", r)
}
