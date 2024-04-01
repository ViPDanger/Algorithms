package Algoritms

import (
	"fmt"
	"math"
)

func Min2MaxArray(arr []int) (int, int) {
	max := arr[0]
	min := max
	for i := 0; i <= len(arr)-1; i++ {
		if arr[max] < arr[i] {
			max = i
		}
		if arr[min] > arr[i] {
			min = i
		}
	}
	return min, max

}

func Test4() {
	var n, k int
	fmt.Scan(&n, k)
	a := make([]int, n)

	hash := make(map[int]int)

	// В процессе записи сразу запишем возможные операции в hash таблицу
	for i := 0; i < n; i++ {
		fmt.Scan(&(a[i]))
		for j := 9; j >= 0; j-- {
			if 90 > a[i]/int(math.Pow10(j-1)) && (a[i]/int(math.Pow10(j-1)) > 9) {
				hash[j] = a[i]
			}
		}
	}
	// Ну и
	//var ans int64
	for i := 0; i <= k; i++ {
		//	sort.Ints()
	}
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
