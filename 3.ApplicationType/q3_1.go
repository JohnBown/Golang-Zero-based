package zerobased

import (
	"fmt"
)

// Q3_1 编写一个程序，写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。
func Q3_1() {
	var a [16]int

	for i := 0; i <= 15; i++ {
		a[i] = i
	}

	fmt.Println(a)
}
