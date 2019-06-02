package zerobased

import "fmt"

// Q3_10 编写一个程序，定义1个map，并为其赋值，再定义一个切片，
// 然后遍历map所有成员，将该map所有成员的value添加到切片中。
func Q3_10() {
	weeks := map[int]string{
		1: "Mon",
		2: "Tue",
		3: "Wed",
		4: "Thu",
		5: "Fri",
		6: "Sat",
		7: "Sun",
	}
	s := []string{}

	for _, v := range weeks {
		s = append(s, v)
	}

	fmt.Println(s)
}
