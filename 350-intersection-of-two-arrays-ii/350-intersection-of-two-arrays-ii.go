func intersect(nums1 []int, nums2 []int) []int {
    set := make(map[int]int)
    for _, v := range nums1 {
        set[v]++
    }
    
    res := make([]int, 0)
    for _, v := range nums2 {
        if set[v] > 0 {
            res = append(res, v)
            set[v]--
        }
    }
    
    return res
}