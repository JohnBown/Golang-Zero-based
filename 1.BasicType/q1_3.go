package zerobased

import (
	"fmt"
)

// Q1_3 编写一个程序，定义2个整型变量，将第一个变量除以第二个变量的值（浮点数型）
// 输出到屏幕上。
func Q1_3() {
	var a, b int
	a, b = 10, 3
	expr := float32(a) / float32(b)

	fmt.Printf("%d / %d = %f\n", a, b, expr)
}
