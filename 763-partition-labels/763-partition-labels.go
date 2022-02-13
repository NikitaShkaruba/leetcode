func partitionLabels(S string) []int {
    s := S
    dict := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        dict[s[i]] = i
    }
    
    var res []int
    
    i := 0
    j := dict[s[0]]
    cnt := 1
    for {
        if i == j {
            res = append(res, cnt)
            i++
            if i >= len(s) {
                break
            }
            j = dict[s[i]]
            cnt = 1
        } else {
            j = max(j, dict[s[i]])
            cnt++
            i++
            if i >= len(s) {
                break
            }
        }
    }
    
    return res
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
