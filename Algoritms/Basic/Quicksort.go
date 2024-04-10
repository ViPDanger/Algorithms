package basic

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	wall := 0
	pivot := a[len(a)-1]
	for current := 0; current < len(a); current++ {
		if a[current] < pivot {
			a[wall], a[current] = a[current], a[wall]
			wall++
		}
	}
	a[wall], a[len(a)-1] = a[len(a)-1], a[wall]
	return append(QuickSort(a[:wall]), QuickSort(a[wall:])...)
}

func QuickSortStr(str string) string {
	return string(quicksortrune([]rune(str)))
}
func quicksortrune(a []rune) []rune {
	if len(a) < 2 {
		return a
	}
	wall := 0
	pivot := a[len(a)-1]
	for current := 0; current < len(a); current++ {
		if a[current] < pivot {
			a[wall], a[current] = a[current], a[wall]
			wall++
		}
	}
	a[wall], a[len(a)-1] = a[len(a)-1], a[wall]
	return append(quicksortrune(a[:wall]), quicksortrune(a[wall:])...)
}
