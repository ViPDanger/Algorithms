package Output

import (
	"fmt"
	"math/rand"
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

func Create_str(s *string, min int, max int, count int) {

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		m := rand.Intn((max - min)) + min

		(*s) = (*s) + " " + strconv.Itoa(m)
		time.Sleep(time.Nanosecond)
	}
}

func Println(a ...any) (n int, err error) {
	n, err = fmt.Println(a)
	return n, err
}
