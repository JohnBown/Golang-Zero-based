package zerobased

import (
	"fmt"
)

func reverse(s *string) {
	b := []byte(*s)

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	*s = string(b)
}

// Q3_4 编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”，
// 接受键盘输入字符串，然后将反转后的字符串打印到屏幕上；
func Q3_4() {
	s := "Google"
	reverse(&s)
	fmt.Println(s)
}
