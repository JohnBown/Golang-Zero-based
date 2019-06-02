package zerobased

import (
	"fmt"
	"strconv"
)

// Q3_8 编写一个程序，创建一个 map，key类型为字符串，value类型为字符串 ，
// 并未其8个成员赋值，检查是否有成员的value为“hello”，如果有，在屏幕上打印”hello exist”，
// 如果没有，在屏幕上打印“hello not exist”
func Q3_8() {
	m := make(map[string]string)

	for i := 1; i <= 8; i++ {
		var tmp string
		fmt.Scanf("%s", &tmp)

		m[strconv.Itoa(i)] = tmp
	}

	flag := "hello not exist"
	for _, v := range m {
		if v == "hello" {
			flag = "hello exist"
		}
	}

	fmt.Println(flag)
}
