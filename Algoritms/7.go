package Algoritms

import fmt "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"

type stud struct {
	b bool   // Ссылаются ли на этого ученика
	a uint32 // Номер ученика, на которого ссылаются
}

func Test7() {

	var n, i, j, count, no_count, a uint32
	fmt.Scan(&n)
	fmt.Println(n)
	count = 0
	no_count = 0
	ai := make([]stud, n+1)
	// Вычисление единственного
	for i = 1; i <= n; i++ {
		fmt.Scan(&a)
		fmt.Println(a)
		(ai[i]).a = a
		// Если на это число уже ссылались - мы записываем номер
		if ai[a].b {
			if count == 0 {
				count = a
			} else {
				println(-1, -1)
				return
			}

		} else {
			ai[a].b = true
		}
		// Если на i не ссылались - мы записываем его в нессылающиеся
		if !ai[i].b {
			no_count = i
		}
	}
	if count == 0 || no_count == 0 {
		fmt.Println(-1, -1)
		return
	}
	ai[count].a = no_count

	i = 1
	for j = 1; true; j++ {
		i = ai[i].a
		if j >= n {
			break
		}
	}
	if i != 1 {
		fmt.Println(-1, -1)
	} else {
		fmt.Println(count, no_count)
	}
}
