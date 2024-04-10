package main

import (
	"fmt"
	"time"

	tinkoff "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Tinkoff"
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
	var s string
	output.Create_str(&s, 1, 999, 10)
	output.Str = s
	fmt.Println(s)
	time1 := time.Now()
	//v := "asdasdkhfuygqywe819361jkfdngkajhijahsd71236nvosof"
	//fmt.Println(s)
	output.PrintMemUsage()
	tinkoff.Test3()
	
	fmt.Printf("%f %s \n", float64(time.Since(time1).Nanoseconds())/1000000000., " s")
	output.PrintMemUsage()
	//fmt.Println(s)

}
