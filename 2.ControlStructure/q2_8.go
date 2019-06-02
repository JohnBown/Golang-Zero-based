package zerobased

import (
	"fmt"
)

// Q2_8 编写一个程序，使用for循环，接收键盘输入，当输入的值是一个大于5的数字时，
// 跳出循环，并将输入的大于5的值输出到屏幕上；
func Q2_8() {
	var a int

	for {
		fmt.Scanf("%d", &a)
		if a > 5 {
			fmt.Println(a)
			break
		}
	}
}
