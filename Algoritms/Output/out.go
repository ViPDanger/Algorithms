package Output

import (
	"fmt"
	"strings"
)

var Str string

// var filename = "Algoritms/Output/buffer.txt

func Scan(a ...any) (n int, err error) {
	for _, i := range a {
		reader := strings.NewReader(Str)
		n, err = fmt.Fscan(reader, i)
		fmt.Fscan(reader, i)
		Str = strings.Join(strings.Split(Str, " ")[:], " ")
		fmt.Println(Str)

	}

	return n, err
}

func Println(a ...any) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}
