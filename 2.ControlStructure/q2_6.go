package zerobased

import (
	"fmt"
)

// Q2_6 编写一个程序，接收键盘输入一个代表月份的数字，然后
// 返回所代表月份所在季节，将季节名称输出到屏幕上。
func Q2_6() {
	var a int
	fmt.Scanf("%d", &a)

	switch {
	case 3 <= a && a <= 5:
		fmt.Println("spring")
	case 6 <= a && a <= 8:
		fmt.Println("summer")
	case 9 <= a && a <= 11:
		fmt.Println("autumn")
	case a == 12 || a == 1 || a == 2:
		fmt.Println("winter")
	default:
		fmt.Println("failing")
	}
}
