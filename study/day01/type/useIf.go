package main

import "fmt"

// if逻辑控制的操作
func main6() {

	// if在判断逻辑时候还能进行变量的初始化
	if result := f1(); result == 1 {
		fmt.Println("result = 1")
	} else {
		fmt.Println("result != 1")
	}
}

func f1() (result int) {
	result = 2
	return result
}
