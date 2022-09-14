package main

import "fmt"

// 自定义函数
func main() {
	result := MyFunc1(10, 20)
	fmt.Println("主函数的调用结果：", result)

	r1, r2 := MyFunc2(20, 10)
	fmt.Println("两个数相加的结果是：", r1, "两个数相减的结果是：", r2)

	r3 := MyFunc3(10, 20, 30, 40)
	fmt.Println("多个数相加的结果是：", r3)
}

func MyFunc1(num1, num2 int) int {
	fmt.Println(num1 + num2)
	return num1 + num2
}

func MyFunc2(num1, num2 int) (plus, sub int) {
	plus = num1 + num2
	sub = num1 - num2
	return plus, sub
}

// 不定参数一定要放到最后一个参数,不然会分不清到底是哪个参数
func MyFunc3(num int, data ...int) int {
	var result int
	for _, ele := range data {
		result += ele
	}
	fmt.Println("传入的num为：", num)
	result -= num
	return result
}

// 使用不定参数给其他函数传递参数
func MyFunc4(data ...int) {
	MyFunc3(data[:1]...)
}
