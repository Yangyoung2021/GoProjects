package main

import (
	"fmt"
	"os"
)

func main() {
	var sum string
	var s1 = []string{"aaa", "bbb", "ccc"}
	sep := ", "
	for i := 0; i < len(os.Args); i++ {
		sum += os.Args[i] + sep
	}
	fmt.Println("命令行参数是：", sum)
	sum = ""
	for k := 0; k < len(s1); k++ {
		sum += s1[k] + sep
	}
	fmt.Println("切片的值为：", sum)
}
