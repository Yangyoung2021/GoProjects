// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个map集合
	counts := make(map[string]int)
	// 把保存输入的扫描仪的地址赋值为input对象
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("input's type is %T\n", input)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
