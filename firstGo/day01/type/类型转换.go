package main

import "fmt"

func main3() {
	var a = 100
	var c1 = rune(a)
	// go语言中的类型转换是使用类型名称然后使用括号包裹要转换的变量
	// 这个和Java中的类型转换刚好相反
	// go中的字符类型占用4个字节，不会显示出定义的rune类型，而是展示int32
	fmt.Printf("c1 type is %T and a is %c\n", c1, c1)
	var b byte = 1
	// byte的类型打印出来时int8，一个字节
	fmt.Printf("b type is %T\n", b)

	a = int(c1)
	fmt.Printf("a = %d\n", a)
}
