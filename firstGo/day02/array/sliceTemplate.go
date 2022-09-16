package main

import (
	"fmt"
)

func main2() {
	var s1 = []string{"aaa", "BB", "111", "NNN"}
	// 切片遍历
	for index, data := range s1 {
		fmt.Printf("s1[%d] = %s\n", index, data)
	}
	// 切片定义范围规则，左闭右开
	fmt.Println("s1[2:3]的长度 ", len(s1[2:3]))
	fmt.Println(s1)
}
