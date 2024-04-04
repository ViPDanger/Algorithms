package Tinkoff

import (
	"math"

	fmt "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
)

func Test4() {
	var n, a int
	var j uint8
	var i, k uint
	var ans uint64
	fmt.Scan(&n, &k)
	a_op := make(map[uint8][]uint, 10)
	for j = 0; j < 10; j++ {
		a_op[j] = make([]uint, 10)

	}

	// В процессе записи сразу запишем количество каждой возможной операции в таблицу
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		for j = 0; a > 0; j++ {
			a_op[j][9-a%10] += 1
			a = a / 10

		}
	}

	for j = 9; j < math.MaxUint8 && k > 0; j-- {
		for i = 8; i < math.MaxUint8 && k > 0; i-- {

			if a_op[j][i] != 0 {
				if k < a_op[j][i] {
					a_op[j][i] = k
				}
				ans += uint64(i*a_op[j][i]) * uint64(math.Pow10(int(j)))
				k -= a_op[j][i]

			}

		}
	}
	fmt.Println(ans)
}
