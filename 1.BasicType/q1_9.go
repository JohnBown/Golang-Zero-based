// 编写一个程序，定义一个圆周率常量Pi，值为3.1415926，接收键盘输入圆的半径，
// 求圆的周长，将求得的值输出到屏幕上。
package zerobased

import (
	"fmt"
)

func Q1_9(){
	const Pi float64 = 3.1415926

	var r float64
	fmt.Scanln(&r)

	circum := Pi * r / float64(2)

	fmt.Printf("When the radius of the circle is %.2f, the circumference is %.2f.\n", r, circum)
}