package main

import "fmt"

func main() {
	var a, b, c float64;
	a = 1.69
	b = 1.7
	c = a * b//正确是2.873
	fmt.Print(c)//实际是2.8729999999999998
}
