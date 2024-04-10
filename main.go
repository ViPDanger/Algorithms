package main

import (
	"fmt"
	"time"

	"github.com/ViPDanger/AlgoritmicProblems/Algoritms/basic"
	l "github.com/ViPDanger/AlgoritmicProblems/Algoritms/letcode"
	output "github.com/ViPDanger/AlgoritmicProblems/Algoritms/output"
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
	output.Create_int(&s, 1, 999, 5000000)
	//fmt.Println(s)
	time1 := time.Now()
	//v := "asdasdkhfuygqywe819361jkfdngkajhijahsd71236nvosof"
	//fmt.Println(s)
	basic.QuickSort(s)

	fmt.Printf("%f %s \n", float64(time.Since(time1).Nanoseconds())/1000000000., " s")
	output.PrintMemUsage()
	//fmt.Println(s)

}
