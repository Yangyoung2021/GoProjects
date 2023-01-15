package main

import (
	"fmt"
)

func main1() {
	// 声明固定长度的数组
	var arr2 [3]string
	var temp = "good"
	// 为数组赋值
	arr2 = [3]string{temp, "bad", "fantastic"}
	fmt.Println(arr2)
	// 声明固定长度并赋值数组
	var arr1 = [3]int{1, 2, 3}
	// 数组可以直接打印获取其中的数据对象
	fmt.Println(arr1)
	// 声明不定长度的数组并通过赋值确定长度
	var arr = [...]string{"1", "2", "3", "4"}
	fmt.Println(arr)
	var tempPtr = &arr
	changeArr(tempPtr)
	fmt.Println(arr)
	str := "aaa"
	changeString(&str)
	fmt.Println(str)

}

func changeArr(arr *[4]string) {
	// 遍历的同时进行修改
	for i := 0; i < len(*arr); i++ {
		arr[i] = "abc"
	}
	// 只遍历，不修改
	for i, data := range *arr {
		fmt.Printf("arr[%d] = %v\n", i, data)
	}
}

func changeString(str *string) {
	*str = "bbb"
}
