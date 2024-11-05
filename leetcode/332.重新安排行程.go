package leetcode

import "sort"

// [332] 重新安排行程
type Tickets [][]string
func (t Tickets) Len() int {
    return len(t)
}
func (t Tickets) Less(i, j int) bool {
    if t[i][0] < t[j][0] {
		return true
	} else if t[i][0] > t[j][0] {
		return false
	}
	if t[i][1] < t[j][1] {
		return true
	}
	return false
}
func (t Tickets) Swap(i, j int) {
    t[i], t[j] = t[j], t[i]
}

type FromTo struct{
    Used bool
    To string
}
func findItinerary(tickets [][]string) []string {
    ts := Tickets(tickets)
    sort.Sort(ts)
    l := len(ts)
    path := []string{"JFK"}

    fromTos := make(map[string][]*FromTo)
    for _, t := range tickets{
        tos, ok := fromTos[t[0]]
        fromTo :=  &FromTo{
            Used: false,
            To: t[1],
        }
        if !ok {
            fromTos[t[0]] = []*FromTo{fromTo}
        } else {
			fromTos[t[0]] = append(tos, fromTo)
        }
    }
    var dfs func(from string) bool
    dfs = func(from string) bool {
        ll := len(fromTos[from])
        if len(path) == l + 1 {
            return true
        }
        for i := 0; i < ll; i++ {
            fromTo := fromTos[from][i]
            if fromTo.Used {
                continue
            }
            fromTo.Used = true
            path = append(path, fromTo.To)
            if dfs(fromTo.To) {
                return true
            }
            fromTo.Used = false
            path = path[0 : len(path) - 1]
        }
        return false
    }
    dfs("JFK")
    return path
}