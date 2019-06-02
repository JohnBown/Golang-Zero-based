package zerobased

import (
	"fmt"
)

// Q1_1 编写一个程序，定义两个字符串变量，为两个变量赋值后，
// 将两个变量的值相加并输出到屏幕上。
func Q1_1() {
	var hello, world string

	hello = "Hello,"
	world = "World"
	splice := hello + world

	fmt.Println(hello, world, " = ", splice)
}
