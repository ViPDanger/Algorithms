package Algoritms

import (
	"fmt"
)

// Функция возвращает индексы Min Max

func MinMaxArray(arr []int) (int, int) {
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

func Test3() {
	var n, t, nt int

	// Считывание переменных
	fmt.Scanln(&n, &t)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&(arr[i]))
	}
	fmt.Scan(&nt)

	nt = nt - 1
	min, max := MinMaxArray(arr)
	min = arr[min]
	max = arr[max]

	// Вывод данных
	// Если мы не успеваем ни с минимума, ни с максиумума, то будем идти от точки nt.
	if (max-arr[nt] > t) && (arr[nt]-min > t) {
		fmt.Println(max - min)
	} else if (2*max - min - arr[nt]) < (max - 2*min + arr[nt]) {
		fmt.Println(2*max - min - arr[nt])
	} else {
		fmt.Println(max - 2*min + arr[nt])
	}
}
