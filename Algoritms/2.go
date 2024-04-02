package Algoritms

import (
	"fmt"
	"math"
)

func Test2() {
	var a int32
	fmt.Scan(&a)
	fmt.Println(math.Ceil(math.Log2(float64(a))))
}
