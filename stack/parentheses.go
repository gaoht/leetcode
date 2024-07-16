package stack

// IsValid 20 括号匹配
func IsValid(s string) bool {
	b := []byte(s)
	t := NewStack()
	
	m := make(map[byte]byte)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	for _, v := range b {
		if v == ')' || v == ']' || v == '}' {
			if t.IsEmpty() {
				return false
			}
			l := byte(t.Pop())
			if m[v] != l {
				return false
			}
		} else {
			t.Push(int(v))
		}
	}
	return t.IsEmpty()
}