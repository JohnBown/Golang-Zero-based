// 编写一个程序，定义一个结构体，结构体成员包含用户名、身份证号码、余额(浮点数类型)，将结构体信息输出到屏幕上。
package zerobased

import (
	"fmt"
)

func Q1_4(){
	type Account struct {
		ID 		string
		Name	string
		Balance	float32
	}

	account := Account{ID: "201903110001", Name:"LaoWang", Balance:32.5}
	fmt.Printf("%#v\n", account)
}