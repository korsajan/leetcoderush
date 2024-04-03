package stackqueue

func decodeString(s string) string {
	stack := make([]byte, 0)
	var (
		str, num string
	)
	for _, ch := range []byte(s) {
		if ch != ']' {
			stack = append(stack, ch)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] != '[' {
			str = string(stack[len(stack)-1]) + str
			stack = stack[:len(stack)-1]
		}
		stack = stack[:len(stack)-1]

		for len(stack) > 0 && isDigit(stack[len(stack)-1]) {
			num = string(stack[len(stack)-1]) + num
			stack = stack[:len(stack)-1]
		}

		for i := 0; i < atoi(num); i++ {
			stack = append(stack, str...)
		}
		str, num = "", ""
	}

	return string(stack)
}

func isValid(s string) bool {
	var brackets = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	bb := make([]byte, 0)
	for _, b := range []byte(s) {
		if b == '{' || b == '[' || b == '(' {
			bb = append(bb, b)
		} else if len(bb) > 0 && bb[len(bb)-1] == brackets[b] {
			bb = bb[:len(bb)-1]
		} else {
			return false
		}
	}
	return len(bb) == 0
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func atoi(s string) int {
	var (
		n    int
		sign bool
	)
	for _, b := range []byte(s) {
		if b == '-' {
			sign = true
			continue
		}
		if b > '9' {
			break
		}
		n = n*10 + int(b) - '0'
	}
	if sign {
		return n * -1
	}
	return n
}
