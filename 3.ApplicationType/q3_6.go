package zerobased

import (
	"fmt"
)

// Q3_6 编写一个程序，创建一个 map 来保存每周 7 天的名字，然后将它们打印出来。
func Q3_6() {
	weeks := map[int]string{
		1: "Mon",
		2: "Tue",
		3: "Wed",
		4: "Thu",
		5: "Fri",
		6: "Sat",
		7: "Sun",
	}

	for i := 1; i <= 7; i++ {
		fmt.Printf("%d\t%s\n", i, weeks[i])
	}
}
