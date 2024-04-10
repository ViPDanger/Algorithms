package basic

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
