package Output

import (
	"fmt"
	"os"
)

func Scan(a ...any) (n int, err error) {
	const filename = "Algoritms/Output/buffer.txt"
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	n, err = fmt.Fscan(file, a)
	defer file.Close()

	return n, err
}

func PrintBuff(a ...any) (n int, err error) {
	const filename = "Algoritms/Output/buffer.txt"
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	n, err = fmt.Fprint(file, a)
	defer file.Close()

	return n, err
}

func Println(a ...any) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}
