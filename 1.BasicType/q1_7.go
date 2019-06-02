package zerobased

import (
	"fmt"
)

// Q1_7 编写一个程序，定义一个一维数组，通过指针为其赋值，并将数组内容输出到屏幕上。
func Q1_7() {
	var arr []int
	p := &arr
	*p = []int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr)
}
