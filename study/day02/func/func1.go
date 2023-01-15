package main

import "fmt"

// 自定义函数
func main1() {
	result := MyFunc1(10, 20)
	fmt.Println("主函数的调用结果：", result)

	r1, r2 := MyFunc2(20, 10)
	fmt.Println("两个数相加的结果是：", r1, "两个数相减的结果是：", r2)

	MyFunc4(10, 20, 30, 40)
	fmt.Println("多个数相加的结果是：")

	fmt.Println()

	min, max := Myfunc5(10, 23, 84)
	fmt.Println("最大值是：", max, ", 最小值是：", min)
}

func MyFunc1(num1, num2 int) int {
	fmt.Println(num1 + num2)
	return num1 + num2
}

func MyFunc2(num1, num2 int) (r1, r2 int) {
	r1 = num1 + num2
	r2 = num1 - num2
	return r1, r2
}

// 不定参数一定要放到最后一个参数,不然会分不清到底是哪个参数
func MyFunc3(num int, data ...int) int {
	var result int
	fmt.Printf("获得的参数是：")
	for index, ele := range data {
		fmt.Printf("参数[%d]为%d\t", index, ele)
		result += ele
	}
	fmt.Println()
	fmt.Println("传入的num为: ", num)
	result -= num
	return result
}

// 使用不定参数给其他函数传递参数
func MyFunc4(data ...int) {
	// 使用[]进行不定参数的传递时，[]中的数字表示对应的索引值，
	// :在前面就表示索引值之前，后面就表示索引值之后
	// 并且索引在:左边就会包括索引值，在:右边就会不会包括索引值（在:左闭右开）
	result := MyFunc3(data[0], data[1:3]...)
	fmt.Println("结果是：", result)
}

func Myfunc5(num1, num2, num3 int) (min, max int) {
	if num1 > num2 {
		if num1 > num3 {
			max = num1
			if num3 > num2 {
				min = num2
			} else {
				min = num3
			}
		} else {
			max = num3
			min = num2
		}
	} else {
		if num2 > num3 {
			max = num2
			if num3 > num1 {
				min = num1
			} else {
				min = num3
			}
		} else {
			max = num3
			min = num1
		}
	}

	return min, max
}
