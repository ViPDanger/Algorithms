package Tinkoff

import fmt "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"

func Test3() {
	var n, t, nt int
	// Считывание переменных
	fmt.Scan(&n, &t)
	arr := make([]int, n)
	fmt.Scan(&(arr[0]))
	min := arr[0]
	max := min

	// При считывании массива сразу найдём в нём минимум и максимум
	for i := 1; i <= n; i++ {

		fmt.Scan(&(arr[i]))
		if max < arr[i] {
			max = arr[i]
		} else if min > arr[i] {
			min = arr[i]
		}
	}
	fmt.Scan(&nt)

	nt = nt - 1
	// Вывод данных
	// Если мы не успеваем ни с минимума, ни с максиумума, то будем идти от точки nt.
	if (max-arr[nt] <= t) || (arr[nt]-min <= t) {
		fmt.Println(max - min)
	} else if (max - arr[nt]) < (arr[nt] - min) {
		fmt.Println(2*max - min - arr[nt])
	} else {
		fmt.Println(max - 2*min + arr[nt])
	}
}
