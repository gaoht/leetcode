package strings

// ReverseWords 151
func ReverseWords1(s string) string {
    b := []byte(s)
    ret :=  make([]byte, 0)
    
    for l, i := len(b), len(b) - 1; i >= 0; i--{
        if b[i] != ' ' && i != 0{
            continue
        }
		// 连续空格跳过
        if b[i] == ' ' && i+1 < len(b) && b[i + 1] == ' ' {
            l = i
            continue
        }
        if b[i] == ' ' {
            if len(ret) > 0 && ret[len(ret) - 1] != ' ' {
                ret = append(ret, ' ')
            }
            ret = append(ret, b[i+1:l]...)
        } else {
			// 处理开始
            if len(ret) > 0 && ret[len(ret) - 1] != ' ' {
                ret = append(ret, ' ')
            }
            ret = append(ret, b[i:l]...)
        }
        l = i
    }
    return string(ret)
}


func ReverseWords(s string) string {
    b := []byte(s)
	b = removeSpace(b)
    reverseBytes(b, 0, len(b) - 1)
	last := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' || i == len(b) - 1{
			e := i - 1
			if i == len(b) - 1 {
				e = i
			}
			reverseBytes(b, last, e)
			last = i + 1
		}
	}
    return string(b)
}

func removeSpace(b []byte)[] byte{
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' '{
			if slow > 0 {
				b[slow] = ' '
				slow ++
			}
			for i < len(b) && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	return b[0:slow]
}

func reverseBytes(b []byte, s int, e int) []byte{
	left := s
	right := e
	for left < right {
		tmp := b[right]
		b[right] = b[left]
		b[left] = tmp
		right--
		left++
	}
	return b
}