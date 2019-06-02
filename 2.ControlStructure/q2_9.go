package zerobased

import (
	"fmt"
)

// Q2_9 编写一个程序，接收键盘输入一个数字hour，如果是0<hour<12，输出am+hour值
// 到屏幕上，如果12<hour<24,输出pm+hour值到屏幕上；
func Q2_9() {
	var hour int

	fmt.Scanf("%d", &hour)

	if 0 < hour && hour < 12 {
		fmt.Println("am", hour)
	} else if 12 < hour && hour < 24 {
		fmt.Println("pm", hour)
	}
}
