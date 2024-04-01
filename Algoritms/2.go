package Algoritms

import (
	"fmt"
)

func Test2() {
	var a, i int
	fmt.Scan(&a)
	for i = 0; a > 1; i++ {
		a = a/2 + a%2
	}
	fmt.Println(i)
}
