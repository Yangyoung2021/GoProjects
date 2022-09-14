package main

import "fmt"

func main3() {

	var c1 = Plus
	result := c1(10, 20)

	fmt.Println("result = ", result)

	r1 := Judge(20, 10, Sub)
	fmt.Println("r1结果是：", r1)

}

func Plus(num1, num2 int) (result int) {
	result = num1 + num2
	return result
}

func Sub(num1, num2 int) (result int) {
	result = num1 - num2
	return result
}

func Judge(n1 int, n2 int, f1 Calc) int {
	return f1(n1, n2)
}

// 定义一种函数类型
type Calc func(int, int) int
