package main

import "fmt"

// switch控制语句的使用
func main7() {
	var a = 100
	/**
	switch相对其他语言来说，多了可以在要判断变量的同时声明一个变量，与此同时还能在switch
	的时候不输入条件，然后case块中直接使用条件语句进行控制
	*/
	switch b := 20; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3:
		fmt.Println("b = 4")
	case 4:
		fmt.Println("b = 5")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b是其他值")
	}

	fmt.Println("开始进行判断值为空的演示")

	switch {
	case a > 100:
		fmt.Println("a > 100")
	case a == 100:
		fmt.Println("a = 100")
	case a < 100:
		fmt.Println("a < 100")
	}
	fmt.Println()
	sum := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
			sum += i * j
		}
		fmt.Println()
	}
	fmt.Println("sum = ", sum)

}
