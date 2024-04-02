package main

import (
	"fmt"
	"os"
	"strconv"

	alg "github.com/ViPDanger/AlgoritmicProblems/Algoritms"
)

func main() {

	alg.Test7()
	fmt.Println("Вупа вупа")
	for i := 3; i < 100; i++ {
		os.Stdin.Write([]byte(strconv.Itoa(i)))
	}

}
