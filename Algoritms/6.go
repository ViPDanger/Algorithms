package Algoritms

import "fmt"

func Test6() {
	var n, a, m1, m2, i int64
	fmt.Scan(&n)
	m1 = -1
	m2 = -1
	for i = 1; i <= n; i++ {
		fmt.Scan(&a)
		if i%2 != a%2 {
			if m2 == -1 {
				m2 = m1
				m1 = i
			} else {
				fmt.Println(-1, -1)
				return
			}
		}
	}
	fmt.Println(m2, m1)
}
