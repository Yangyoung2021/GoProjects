package main

import "fmt"

func main() {
	var str = "abcdefg"
	// 循环的遍历方式一，使用索引遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\t", i, str[i])
	}
	fmt.Println()
	// 遍历方式二，使用增强for循环
	for index, data := range str {
		fmt.Printf("str[%d] = %c\t", index, data)
	}

}
