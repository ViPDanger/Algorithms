package main

import (
	"strconv"
	"time"

	alg "github.com/ViPDanger/AlgoritmicProblems/Algoritms"
	buf "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
)

func main() {
	str := strconv.Itoa(1)
	buf.Create_str(&str, 1, 500, 500)
	buf.Str = str
	time1 := time.Now()
	alg.Test2()
	println(time.Since(time1).String())

}
