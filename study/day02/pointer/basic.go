package main

import "fmt"

// &表示取地址，用在变量前
// *表示两种意思，如果后面接的是类型表明前面的变量为指针类型，如果后面接的是变量代表取变量中的值
func main() {
	// 声明int类型的指针ip
	var ip *int
	// 声明float类型的指针fp
	var fp *float32
	// 声明string类型的指针sp
	var sp *string
	a := 10
	var b float32 = 100.0
	// 指针的赋值只能通过已声明的对象进行取地址操作
	ip = &a
	fp = &b
	// 打印指针指向地址保存的值
	fmt.Println("ip指针中保存的值为: ", ip)
	fmt.Println("ip指针中保存的值所对应地址保存的值为: ", *ip)
	fmt.Println("保存ip指针地址的值为: ", &ip)
	fmt.Println("fp指针中保存的值为: ", fp)
	fmt.Println("fp指针中保存的值所对应地址保存的值为: ", *fp)
	fmt.Println("保存fp指针地址的值为: ", &fp)
	// 为赋值的指针默认值为null
	fmt.Println("sp指针中保存的值为: ", sp)
	// 如果对值为null的指针进行取值操作就会触发异常（空指针）
	fmt.Println("sp指针中保存的值所对应地址保存的值为: ", *sp)
	fmt.Println("保存sp指针地址的值为: ", &sp)
}
