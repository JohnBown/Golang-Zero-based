package zerobased

import (
	"fmt"
)

// Q2_2 编写一个程序，接收键盘输入，如果输入的字符是password，在屏幕上
// 输出ok，如果不是，在屏幕上输出“please enter password”。
func Q2_2() {
	var s string
	fmt.Scanln(&s)

	if s == "password" {
		fmt.Println("ok")
	} else {
		fmt.Println("please enter password")
	}
}
