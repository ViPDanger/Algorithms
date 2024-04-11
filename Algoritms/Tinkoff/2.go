package tinkoff

import (
	"math"

	fmt "github.com/ViPDanger/AlgoritmicProblems/algoritms/output"
)

func Test2() {
	var a uint64
	fmt.Scan(&a)
	fmt.Println(math.Ceil(math.Log2(float64(a))))
}
