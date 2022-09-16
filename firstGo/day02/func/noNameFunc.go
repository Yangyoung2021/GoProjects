package main

import "fmt"

// 创建匿名函数并调用
// 匿名函数可以获取所在范围的函数的变量
func main4() {

	num := 100
	str := "abc"

	f1 := func() {
		fmt.Println(num, str)
		fmt.Println("匿名函数的对象被调用")
	}

	f1()

	func() {
		fmt.Print("定义匿名函数并调用")
	}()

}
