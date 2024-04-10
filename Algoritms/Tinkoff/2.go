package tinkoff

import (
	"math"

	fmt "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
)

func Test2() {
	var a uint32
	fmt.Scan(&a)
	fmt.Println(math.Ceil(math.Log2(float64(a))))
}
