package Output

import "fmt"

func Scan(a interface{}) (n int, err error) {
	n, err = fmt.Scan(a)
	return n, err
}

func Println(a interface{}) (n int, err error) {
	fmt.Println(a)
	return n, err
}
