package rand

import (
	"fmt"
)

func f4() {
	var place string = "shenzhen"
	var num = 100
	var a, b, c int = 10, 20, 30
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// Println方法可以使用,进行拼接，也能使用+进行拼接
	fmt.Println("工作地点是：", place)
	fmt.Printf("num的值是: %d", num)

}
