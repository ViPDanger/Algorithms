package output

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var Str string

func Scan(a ...any) (n int, err error) {
	for _, i := range a {
		n, err = fmt.Fscan(strings.NewReader(Str), i)
		Str = strings.Join(strings.Split(Str, " ")[1:], " ")

	}

	return n, err
}

func Scanf(f string, a ...any) (n int, err error) {
	for _, i := range a {
		n, err = fmt.Fscanf(strings.NewReader(Str), f, i)
		Str = strings.Join(strings.Split(Str, " ")[1:], " ")

	}

	return n, err
}

func Create_str(s *string, min int, max int, count int) {

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		m := rand.Intn((max - min)) + min

		(*s) = (*s) + " " + strconv.Itoa(m)
		time.Sleep(time.Nanosecond)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: /pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func Create_int(s *[]int, min int, max int, count int) {

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		*s = append(*s, rand.Intn((max-min))+min)
		//time.Sleep(time.Nanosecond)
	}
}

func Println(a ...any) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}
