package tinkoff

import fmt "github.com/ViPDanger/AlgoritmicProblems/algoritms/output"

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	wall := 0
	pivot := a[len(a)-1]
	for current := 0; current < len(a)-1; current++ {
		if a[current] <= pivot {
			a[wall], a[current] = a[current], a[wall]
			wall++
		}
	}
	a[wall], a[len(a)-1] = a[len(a)-1], a[wall]
	QuickSort(a[0:wall])
	QuickSort(a[wall:])
	return a
}

func Test9() int{
	var n,id100 int
	fmt.Scan(&n)
	cost := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&cost[i])
		if cost[i]<=100{
			id100++
		}
	}

cost = QuickSort(cost)
total :=0
	if cost[id100]>100{
	tickets := (n-id100)/2 + (n-id100)%2
	
	for i:=0;i<id100-1;i++{
		total+=cost[i]
	}
	if tickets%2 !=0{
		total+=cost[id100-1]
	}

	return total
	}
	for i:=0;i<n;i++{
		total+=cost[i]
	}
	return total
}

