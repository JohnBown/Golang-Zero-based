package zerobased

import (
	"fmt"
)

// Q1_10 编写一个程序，计算字符串 “"Hello, how is it BlockChain, Hugo?"” 在字符 H 中
// 出现的次数，将结果输出到屏幕上。
func Q1_10() {
	s := "Hello, how is it BlockChain, Hugo?"
	var cnt int

	for _, v2 := range []rune(s) {
		if v2 == 'H' {
			cnt++
		}
	}

	fmt.Println("The number of characters 'H' is: ", cnt)
}
