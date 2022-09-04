package rand

import "fmt"

func f2() {
	var a, b int = 10, 20 // 同时声明多个类型相同的变量
	a, b = b, a           // 交换两个变量的值

	fmt.Printf("a的值为: %d, b的值为: %d\n", a, b) // 格式化输出
}
