package main

import (
	alg "github.com/ViPDanger/AlgoritmicProblems/Algoritms"
	buf "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
)

func main() {

	for i := 2; i < 20; i++ {
		buf.PrintBuff(i)
	}
	alg.Test7()
}
