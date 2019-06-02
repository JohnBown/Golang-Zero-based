package zerobased

import (
	"fmt"
)

// Q3_9 编写一个程序，创建一个 map，key类型为整型，value类型为字符串，
// 用来保存每周 7 天的名字，然后删除掉value为星期天的成员，并将删除后的map所有成员打印出来。
func Q3_9() {
	weeks := map[int]string{
		1: "Mon",
		2: "Tue",
		3: "Wed",
		4: "Thu",
		5: "Fri",
		6: "Sat",
		7: "Sun",
	}

	fmt.Println(weeks)

	for k, v := range weeks {
		if v == "Sun" {
			delete(weeks, k)
		}
	}

	fmt.Println(weeks)
}
