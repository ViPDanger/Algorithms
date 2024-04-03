package Output

import (
	"fmt"
	"os"
	"strings"
)

var Str string
var filename = "Algoritms/Output/buffer.txt"

func Scan(a ...any) (n int, err error) {
	for _, i := range a {
		reader := strings.NewReader(Str)
		n, err = fmt.Fscan(reader, i)
		fmt.Fscan(reader, i)
		Str = strings.TrimLeft(Str, " ")
		fmt.Println(Str)

	}

	return n, err
}
func WriteData(a ...any) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// сохраняем данные в файл
	fmt.Fprint(file, a)
}

func Println(a ...any) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}
