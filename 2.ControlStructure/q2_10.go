package zerobased

import (
	"fmt"
	"runtime"
)

// Q2_10 编写一个程序，读取当前操作系统类型，输出操作系统信息到屏幕上；
func Q2_10() {
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("NumCPU:", runtime.NumCPU())
}
