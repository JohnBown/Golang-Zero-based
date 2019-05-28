// 编写一个程序，使用for循环，生成10个使用当前时间的纳秒级数字作为种子的随机数(浮点数型)，将生成的随机数输出到屏幕上
package zerobased

import (
	"fmt"
	"time"
	"math/rand"
)

func Q2_5()  {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i:=0; i<10; i++{
		fmt.Println(rand.Float64())
	}
}