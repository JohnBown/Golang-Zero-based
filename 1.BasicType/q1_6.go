// 编写一个程序，定义一个二维数组，为其赋值，并将数组内容依次输出到屏幕上。
package zerobased

import (
	"fmt"
)

func Q1_6(){
	var arr [][]int
	arr = [][]int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(arr)
}