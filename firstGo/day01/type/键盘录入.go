package main

import (
	"fmt"
)

func main2() {
	var input string
	fmt.Println("请输入您要输入的字符串: ")

	// Scanf需要指定类型进行接收，但是Scanf不需要使用类型接收
	fmt.Scanf("%v", &input)

	fmt.Printf("您输入的是类型是: %T\n", input)
	fmt.Println("您输入的是: ", input)
}
