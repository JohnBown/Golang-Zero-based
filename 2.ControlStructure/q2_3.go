// 编写一个程序，使用for循环，实现1+2+3+……..+100，将相加结果输出到屏幕上
package zerobased

import (
	"fmt"
)

func Q2_3()  {
	var sum int

	for i := 1; i <= 100; i++{
		sum += i
	}

	fmt.Println(sum)
}