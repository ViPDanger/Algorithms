package basic

import "github.com/ViPDanger/AlgoritmicProblems/algoritms/output"

func DictionarySort(a []int, k int) []int {
	map_a := make(map[int]int)
	for i := range a {
		map_a[a[i]]++
	}
	j := 0
	for i := 0; i <= k; i++ {
		for map_a[i] > 0 {
			a[j] = i
			j++
			map_a[i]--
		}

	}
	output.PrintMemUsage()
	return a
}
