// 编写一个程序，接收键盘输入作为密码，要求输入两次，如果两次输入的密码一致，屏幕输出OK，
// 如果两次输入密码不一致，屏幕输出password error。
package zerobased

import (
	"fmt"
)

func Q1_8(){
	var pwd, pwdagain string

	fmt.Println("New Password:")
	fmt.Scanln(&pwd)
	fmt.Println("Confirm New Password:")
	fmt.Scanln(&pwdagain)

	if pwd == pwdagain {
		fmt.Println("OK")
	} else {
		fmt.Println("password error")
	}
}