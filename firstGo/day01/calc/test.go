package calc

import(
	"fmt"
)

func main() {
	r1 := Add(10, 20)

    fmt.Println("10 + 20 等于", r1)

    r2 := Sub(10, 20)

    fmt.Println("10 - 20 等于", r2)

    r3 := Multiply(10, 20)

    fmt.Println("10 * 20 等于", r3)

    r4 := Divide(10, 20)

    fmt.Println("10 / 20 等于", r4)
}