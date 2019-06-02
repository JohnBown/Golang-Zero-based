package zerobased

import "fmt"

func insertInts(dest []int, src []int, pos int) []int {
	var tmp []int
	tmp = make([]int, len(dest)+len(src))
	pos--

	for i := range tmp {
		if i < pos {
			tmp[i] = dest[i]
		} else if i < pos+len(src) {
			tmp[i] = src[i-pos]
		} else {
			tmp[i] = dest[i-len(src)]
		}
	}

	return tmp
}

// Q3_5 编写一个程序，定义两个切片，并为其赋值，将其中一个切片
// 插入到另一个切片的第三个成员处。
func Q3_5() {
	x := []int{1, 2, 3, 4, 5, 6}
	y := []int{7, 8, 9}

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Println(insertInts(x, y, 3))
}
