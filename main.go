package main

import (
	"fmt"
	"time"

	l "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Letcode"
	"github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
	"github.com/ViPDanger/AlgoritmicProblems/Algoritms/basic"
)

func lists() {
	var xln, yln *l.ListNode
	x := make([]l.ListNode, 10)
	y := make([]l.ListNode, 10)
	for i := 9; i >= 0; i-- {
		x[i] = l.ListNode{Val: i, Next: xln}
		xln = &x[i]
		y[i] = l.ListNode{Val: i + 1, Next: yln}
		yln = &y[i]
	}
	fmt.Println(l.MergeTwoLists(xln, yln))
}

func main() {
	s := make([]int, 0)
	Output.Create_int(&s, 1, 99, 10)
	fmt.Println(s)
	time1 := time.Now()
	fmt.Println(basic.QuickSort(s))
	println(time.Since(time1).String())

}
