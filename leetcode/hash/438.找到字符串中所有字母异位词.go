package leetcode
// [438] 找到字符串中所有字母异位词


func findAnagrams(s string, p string) []int {
	n := make([]int, 0)
	l := len(p)
    traced := 0
	hash := buildHash(p)
	for i := 0; i + l - 1 < len(s); {
		j := i + traced
		in := true
        for  j < i + l {
            if _, ok := hash[s[j]]; !ok {
				in = false
                break
            } else {
                hash[s[j]]--
				traced ++
            }
			j++
        }
        if in {
            if isAnagram1(hash) {
                n = append(n, i)
            }
			hash[s[i]]++
			traced--
            i++
        } else {
            i = j + 1
			hash = buildHash(p)
			traced = 0
        }
	}
	return n
}


func buildHash(p string) map[byte]int {
    hash := make(map[byte]int)
    for _, v := range p {
        hash[byte(v)]++
    }
    return hash
}
func isAnagram1(m map[byte]int) bool {
    for _, v := range m {
        if v != 0 {
            return false
        }
    }
    return true
}