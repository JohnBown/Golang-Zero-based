package zerobased

import (
	"fmt"
)

// Q3_3 编写一个函数，要求其接受两个键盘输入参数，原始字符串 str 和分割索引 i，
// 然后打印出两个分割后的字符串。
func Q3_3() {
	var s string
	var index int

	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &index)

	if index <= len(s) {
		fmt.Println(s[:index])
		fmt.Println(s[index:])
	} else {
		fmt.Println("Index out of range.")
	}
}
