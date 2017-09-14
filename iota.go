package main

import "fmt"

func main()  {
	/*const a int = 1;
	fmt.Print(a)*/
	/*const (
		a = 1
		b = "hello"
		c = true
	)*/
	//const (
	//	a = 1
	//	b
	//	c
	//	d = "hello"
	//	e
	//)
	const (
		a = 5
		b = iota *2
		c
		d
		e
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
