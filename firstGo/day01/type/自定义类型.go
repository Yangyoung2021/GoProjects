package main

import "fmt"

// 类型别名可以用来做继承相关的业务
func main5() {
	// 自定义变量类型bigint
	type bigint int64

	var b1 bigint

	fmt.Printf("b1 type is %T\n", b1)

}
