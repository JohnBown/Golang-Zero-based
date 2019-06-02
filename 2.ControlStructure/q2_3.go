package zerobased

import (
	"fmt"
)

// Q2_3 编写一个程序，使用for循环，实现1+2+3+……..+100，将相加结果输出到屏幕上
func Q2_3() {
	var sum int

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println(sum)
}
