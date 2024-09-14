package main

import "fmt"

func main() {
	var a int = 16
	var b float64 = 0.04
	var c string = "Go! Go!"
	var d bool = true

	e := 3
	f := 2.5
	g := "kama"
	h := false

	fmt.Printf("a: %v, Type: %T\n", a, a)
	fmt.Printf("b: %v, Type: %T\n", b, b)
	fmt.Printf("c: %v, Type: %T\n", c, c)
	fmt.Printf("d: %v, Type: %T\n", d, d)
	fmt.Printf("e: %v, Type: %T\n", e, e)
	fmt.Printf("f: %v, Type: %T\n", f, f)
	fmt.Printf("g: %v, Type: %T\n", g, g)
	fmt.Printf("h: %v, Type: %T\n", h, h)
}
