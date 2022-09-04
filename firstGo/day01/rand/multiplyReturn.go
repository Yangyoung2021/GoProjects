package main

import (
	"fmt"
)

// 主函数，程序的入口
func main() {
	var a, b, c int
	a, b, c = GetThree()
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
// 获取三个返回值的函数
func GetThree() (a, b, c int) {
	return 1, 2, 3
}