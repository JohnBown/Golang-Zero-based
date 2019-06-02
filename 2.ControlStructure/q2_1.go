package zerobased

import (
	"fmt"
	"unicode"
)

// Q2_1 编写一个程序，写一个switch结构，接收键盘输入，当输入的是数字时，
// 屏幕输出提示“您输入的是数字”，当输入的是字母时，屏幕输出提示“您输入的是
// 字母”，当输入的是其他字符时，屏幕输出“输入其他字符”。
func Q2_1() {
	var c rune
	fmt.Scanf("%c", &c)

	switch {
	case unicode.IsLetter(c):
		fmt.Println("您输入的是字母")
	case unicode.IsDigit(c):
		fmt.Println("您输入的是数字")
	default:
		fmt.Println("输入其他字符")
	}
}
