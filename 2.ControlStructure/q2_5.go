package zerobased

import (
	"fmt"
	"math/rand"
	"time"
)

// Q2_5 编写一个程序，使用for循环，生成10个使用当前时间的纳秒级数字作为种子的
// 随机数(浮点数型)，将生成的随机数输出到屏幕上
func Q2_5() {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float64())
	}
}
