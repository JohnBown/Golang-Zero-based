package zerobased

import (
	"fmt"
)

// Q2_4 编写一个程序，使用for循环，实现1+2+3+4+……，当相加结果大于1000时，
// 跳出循环，输出相加结果到屏幕上
func Q2_4() {
	var sum int

	for i := 1; ; i++ {
		sum += i
		if sum > 1000 {
			break
		}
	}

	fmt.Println(sum)
}
