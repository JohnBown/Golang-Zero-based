package zerobased

import (
	"fmt"
)

// Q2_7 编写一个程序，从 1 打印到 100 ，但是每当遇到 3 的倍数时，不打印相应的数字，
// 但打印一次 "Fizz"。遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5
// 的倍数的数，打印 FizzBuzz
func Q2_7() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Printf("Fizz ")
		} else if i%5 == 0 {
			fmt.Printf("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}
