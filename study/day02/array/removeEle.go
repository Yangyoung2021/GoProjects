package main

import (
	"fmt"
)

func main() {
	stringSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	fmt.Println("原始切片", stringSlice)
	stringSlice = removeHeadElement(stringSlice, 1)
	fmt.Println("头部删除一个元素切片", stringSlice)
	stringSlice = removeTailElement(stringSlice, 1)
	fmt.Println("尾部删除一个元素切片", stringSlice)
	stringSlice = removeElementAtIndex(stringSlice, 1, 2)
	fmt.Println("指定索引处删除一个元素切片", stringSlice)
}

// 从头部删除元素
func removeHeadElement(s []string, count int) []string {
	if len(s) < count {
		return nil
	}
	return s[count:]
}

// 从尾部删除元素
func removeTailElement(s []string, count int) []string {
	if len(s) < count {
		return nil
	}
	return s[:len(s)-count]
}

// 从指定位置删除指定个元素
func removeElementAtIndex(s []string, index int, count int) []string {
	if len(s) < count+count {
		return nil
	}
	result := append(s[:index], s[index+count:]...)
	return result
}
