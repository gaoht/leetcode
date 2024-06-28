package hash
// CanConstruct 383
func CanConstruct(ransomNote string, magazine string) bool {
	dict := [26]int{}
    for i := 0; i < len(magazine); i++ {
        dict[magazine[i] - 'a']++
    }
    for i := 0; i < len(ransomNote); i++{
        dict[ransomNote[i] - 'a']--
        if dict[ransomNote[i] - 'a'] < 0 {
            return false
        }
    }
    return true
}