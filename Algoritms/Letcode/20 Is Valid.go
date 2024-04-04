package Letcode

func compare(r rune, b rune) bool {
	if r == b {
		return true
	} else {
		return false
	}

}

func IsValid(s string) bool {
	buf := make([]rune, 0)
	for _, i := range s {
		switch i {

		case ')':
			if len(buf) == 0 {
				return false
			}
			if !compare('(', buf[len(buf)-1]) {
				return false
			}
			buf = append(buf[:len(buf)-1])
		case ']':
			if len(buf) == 0 {
				return false
			}
			if !compare('[', buf[len(buf)-1]) {
				return false
			}
			buf = append(buf[:len(buf)-1])
		case '}':
			if len(buf) == 0 {
				return false
			}
			if !compare('{', buf[len(buf)-1]) {
				return false
			}
			buf = append(buf[:len(buf)-1])
		default:
			buf = append(buf, i)
		}
	}
	return true
}