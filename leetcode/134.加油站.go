package leetcode
func canCompleteCircuit(gas []int, cost []int) int {
    sum := 0
    start := 0
    i := 0
    tried := 0
    c := 0
    stations := len(gas)
    for c < stations && tried < stations {
        i = i % stations
        sum += gas[i] - cost[i]
        i++
        c++
        if sum < 0 {
            sum = 0
            start = i
            tried++
            c = 0
        }
    }
    if c == stations {
        return start
    }
    return -1
}