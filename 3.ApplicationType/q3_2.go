package zerobased

import (
	"fmt"
	"math/rand"
	"time"
)

func min(slice []int) int {
	min := slice[0]
	for _, i := range slice {
		if i < min {
			min = i
		}
	}
	return min
}

// Q3_2 编写一个程序，创建一个int的切片，并为其前20个成员赋值，
// 通过For-range遍历其所有成员，求得最小值，输出到屏幕上；
func Q3_2() {
	slice := make([]int, 20)

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range slice {
		slice[i] = rand.Intn(100)
	}

	fmt.Println(slice)
	fmt.Printf("Min: %d\n", min(slice))
}
