package main

import (
	"time"

	alg "github.com/ViPDanger/AlgoritmicProblems/Algoritms"
	buf "github.com/ViPDanger/AlgoritmicProblems/Algoritms/Output"
)

func main() {
	str := "999 9999"
    buf.Create_str(&str, 1, 10000000, 999)
	print(str)
	buf.Str = str
	time1 := time.Now()
	alg.Test4()
	println(time.Since(time1).String())

}
