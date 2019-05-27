// 编写一个程序，定义3个整型变量，为3个变量赋值后，将3个变量进行不少于5次的四则运算后，将得到的值输出到屏幕上。
package zerobased

import (
	"fmt"
)

func Q1_2(){
	var a, b, c int

	a, b, c = 30, 5, 2
	expr1 := a + b
	expr2 := a - c
	expr3 := b * c
	expr4 := a / b
	expr5 := (c + a/b)*c - a

	fmt.Println("a, b, c: ", a, b, c)
	fmt.Println("a + b: ", expr1)
	fmt.Println("a - c: ", expr2)
	fmt.Println("b * c: ", expr3)
	fmt.Println("a / b: ", expr4)
	fmt.Println("(c + a/b)*c - a: ", expr5)
}