package main

import "fmt"

func swap(k string, d string) (string, string) {
	return d, k
}
func main() {
	str1 := "coco"
	str2 := "loco"
	swap1, swap2 := swap(str1, str2)
	fmt.Println("reverse:", swap1, swap2)
}
