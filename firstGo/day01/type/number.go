package main

import (
	"fmt"
)

func main1() {
	// 同时定义多个变量
	var (
		i  = 10
		f1 = 103.3
		d1 = 103.5
	)
	// 同时定义多个常量
	const (
		aa = 100
		bb = "好家伙"
		cc = 100.84
	)
	// 声明byte类型的数据, uint8表示8位无符号数，int8表示8位有符号数
	// int8类似Java中的byte（-128~127）
	var b int8 = 'd'
	fmt.Printf("b type is %T\t", b)
	fmt.Printf("b = %d\n", b)

	fmt.Printf("aa = %d, bb = %s, cc = %f\n", aa, bb, cc)
	fmt.Printf("i = %d, f1 = %f, d1 = %f\n", i, f1, d1)

	// 声明long类型
	var l int64 = 1038324
	fmt.Printf("l = %d\n", l)

	// 声明float类型
	var f2 float32 = 193.4243
	fmt.Printf("f2 = %f\n", f2)

	// 声明long类型
	var f3 float32 = 193.4243
	fmt.Printf("f3 = %f\n", f3)

	// 声明bool类型
	var b1 = true
	fmt.Printf("b1 type is %T and b1 is %v\n", b1, b1)

	// 声明rune类型(char)的数据
	var r1 rune = 'd'
	fmt.Printf("r1 type is %T and r1 is %c", r1, r1)

}
