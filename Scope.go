package main

import "fmt"

var a int = 1

func main()  {
	fmt.Println(a)
	var a int = 2
	{
		var a int = 3
		fmt.Println(a)
	}
	fmt.Println(a)
}
