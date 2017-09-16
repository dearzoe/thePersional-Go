package main

import "fmt"

var a int = 1

func main()  {
	//%T 查看类型
	//%v 看看结果
	//%b 二进制
	//%d 十进制
	var a uint
	var b uint
	a = 100
	b = 3
	result := a+b//不同类型相加会报错
	fmt.Printf("%T\n %v\n",result,result)
	fmt.Printf("%b %b\n",a << b,a >> b )
	fmt.Printf("%d %d",a << b,a >> b )
}
