package Algoritms

import (
	"fmt"
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
	var n, k, a int
	fmt.Scan(&n, k)
	hash := make(map[uint8]uint8)
	// В процессе записи сразу запишем возможные операции в hash таблицу
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if k > 0 {

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
