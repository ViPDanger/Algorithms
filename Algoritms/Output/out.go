package Output

import (
	"fmt"
	"os"
)

func Scan(a ...any) (n int, err error) {
	const filename = "buffer.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	n, err = fmt.Fscan(file, a)
	defer file.Close()
	
	eturn n, err


func Println(a ..any) (n int, err error) {
	fmt.Println(a
	eturn n, err
}
