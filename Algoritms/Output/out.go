package Output

import (
	"fmt"
	"strings"
)

var Str string

func Scan(a ...any) (n int, err error) {
	reader := strings.NewReader(Str)
	n, err = fmt.Fscanln(reader, a)

	return n, err
}

func Println(a ...any) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}
