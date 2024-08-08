package leetcode
// [49] 字母异位词分组

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	traced := make(map[int]struct{})
	for i := 0; i < len(strs); i++ {
		if _, ok := traced[i]; ok {
			continue
		}
		tmp := make([]string, 0)
		tmp = append(tmp, strs[i])
		traced[i] = struct{}{}
		for j := i + 1; j < len(strs); j++ {
			if _, ok := traced[j]; ok {
				continue
			}
			if isAnagram(strs[i], strs[j]) {
				tmp = append(tmp, strs[j])
				traced[j] = struct{}{}
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}



