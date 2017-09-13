package main

import "fmt"

func main() {
	var str string
	str = "Go语"
	fmt.Println(len(str))
	fmt.Println(str)
	temp := []rune(str)
	fmt.Println(temp)
	fmt.Println(len(temp))
	temp[0] = 'H'
	result := string(temp)
	fmt.Println(result)

}
