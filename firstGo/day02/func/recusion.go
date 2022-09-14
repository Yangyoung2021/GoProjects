package main

import "fmt"

func main2() {
	var input int
	// 使用Scan可以不用指定类型，使用Scanf需要指定类型
	// 使用起来和Println和Printf一样的规则
	// fmt.Scan(&input)
	fmt.Scanf("%d", &input)
	sum := Recusion1(input)
	fmt.Printf("%d以内的正整数之和为：%d", input, sum)
}

func Recusion1(end int) (result int) {
	if end <= 0 {
		return result
	}
	result += end
	end--
	return Recusion1(end) + result
}
