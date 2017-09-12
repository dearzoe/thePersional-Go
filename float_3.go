package main

import "fmt"

func main() {
	var a, b, c float64;
	a = 1.69 * 100
	b = 1.7 * 10
	c = a * b / (100 * 10)//正确是2.873
	fmt.Print(c)//实际是2.873
}
