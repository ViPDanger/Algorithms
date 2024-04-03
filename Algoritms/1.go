package Algoritms

import (
	fmt "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
)

func Test1() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if d > b {
		fmt.Println(a + (d-b)*c)
	} else {
		fmt.Println(a)
	}
}
