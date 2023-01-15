package main

import "fmt"

func main() {
	var str = "abcdefg"

	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for index, data := range str {
		fmt.Printf("str[%d] = %c\n", index, data)
	}

}
