package main

import (
	"fmt"
)

// 声明map类型对象
func main() {
	// 不添加数据式的声明
	counts := map[string]int{}
	for index, data := range counts {
		fmt.Printf("key = %v and value = %v\n", index, data)
	}
	fmt.Println("--------------")
	data := map[int]int{1: 10, 2: 20}
	for i := 1; i <= len(data); i++ {
		fmt.Printf("key = %v and value = %v\n", i, data[i])
	}
	fmt.Println("--------------")
	temp1 := make(map[string]int, 2)
	temp1["hao"] = 18
	temp1["huan"] = 20
	for i := 0; i < len(temp1); i++ {
		fmt.Printf("temp1.key = %v and temp1.value = %v\n", i, temp1)
	}

}
