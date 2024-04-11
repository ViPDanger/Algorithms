package main

import (
	"fmt"
	"time"

	l "github.com/ViPDanger/AlgoritmicProblems/algoritms/letcode"
	output "github.com/ViPDanger/AlgoritmicProblems/algoritms/output"
	"github.com/ViPDanger/AlgoritmicProblems/algoritms/tinkoff"
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
	s := "5 35 40 101 59 63"
	//output.Create_str(&s, 999999998, 999999999, 1000)
	output.Str = s
	fmt.Println(s)
	time1 := time.Now()
	//v := "asdasdkhfuygqywe819361jkfdngkajhijahsd71236nvosof"
	//fmt.Println(s)
	output.PrintMemUsage()
	fmt.Println(tinkoff.Test9())
	
	fmt.Printf("%f %s \n", float64(time.Since(time1).Nanoseconds())/1000000000., " s")
	output.PrintMemUsage()
	//fmt.Println(s)

}
