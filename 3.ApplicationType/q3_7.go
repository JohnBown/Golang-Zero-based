package zerobased

import (
	"fmt"
)

// Q3_7 编写一个程序，创建一个 map，key类型为整型，value类型为字符串 ，
// 并为其8个成员赋值（key分别为1,2,3…8），修改key为1的成员的value为first，
// 然后将所有成员打印出来 。
func Q3_7() {
	m := make(map[int]string)

	for i := 1; i <= 8; i++ {
		var tmp string
		fmt.Scanf("%s", &tmp)

		m[i] = tmp
	}
	m[1] = "first"

	fmt.Println(m)
}
