package main

import (
	"fmt"
	"time"

	l "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Letcode"
)

func main() {
	var xln,yln *l.ListNode
	x := make([]l.ListNode,10)
	y := make([]l.ListNode,10)
	for i:=9;i>=0;i--{
		 x[i] = l.ListNode{Val:i,Next:xln}
		 xln = &x[i]
		 y[i] = l.ListNode{Val:i+1,Next:yln}
		 yln = &y[i]
	} 

	s := []int{1,1,2}
	fmt.Println(s)
	time1 := time.Now()
	fmt.Println(l.MergeTwoLists(xln,yln))
	println(time.Since(time1).String())

}
