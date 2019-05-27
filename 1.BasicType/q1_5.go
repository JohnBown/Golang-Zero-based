// 编写一个程序，定义一个结构体，结构体成员包含用户名、身份证号码、余额(浮点数类型)，将结构体信息以json方式输出到屏幕上。
package zerobased

import (
	"fmt"
	"encoding/json"
)

func Q1_5(){
	type Account struct {
		ID 		string
		Name	string
		Balance	float32
	}

	var accounts = []Account {
		{ID: "201903110001", Name:"LaoWang", Balance:32.5},
		{ID: "201903110002", Name:"LaoLi", Balance:67.8},
		{ID: "201903110003", Name:"LaoZhao", Balance:9.8},
	}
	data, _ := json.MarshalIndent(accounts, "", "    ")

	fmt.Printf("%s\n", data)
}
