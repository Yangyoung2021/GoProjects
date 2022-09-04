package main

import "fmt" // 导入的包必须要使用

func main() {
	const PI = "3.1415926" // 申明了变量，但是常量可以不使用变量必须要使用
	var a int = 0
	a = 100 // 变量赋值
	b := 200
	fmt.Println("b = ", b)
	fmt.Printf("a = %d", a)
}
