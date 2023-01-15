package main

import (
	"fmt"
)

func main4() {
	// 中括号中有数字就是数组，没有数字就是切片
	var s1 = []string{"a", "b", "c"}
	changeSlice(s1)
	// 向切片尾端添加数据
	s1 = append(s1, "hao", "jia")
	// 向切片头部添加数据，只能添加切片，后面的s1需要添加...表示在头部添加
	s1 = append([]string{"1", "2"}, s1...)
	fmt.Println("指定位置插入数据前", s1)
	// 在指定位置添加数据
	s1 = append(s1[:3], append([]string{"t"}, s1[3:]...)...)
	fmt.Println("指定位置插入数据后", s1)
	var s2 = []string{"2"}
	// 在进行拷贝的时候需要目标切片容量能够装下需要拷贝的数据的容量
	copy(s2, s1[3:])
	fmt.Println(s1)
	fmt.Println(s2)
}

// 切片指针作为参数传递的时候只能遍历查看，不能修改
func changeSlice(sl []string) {
	for i, v := range sl {
		sl[i] = "12" + v
	}
}
