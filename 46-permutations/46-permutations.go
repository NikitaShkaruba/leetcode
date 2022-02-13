func permute(nums []int) [][]int {
    if len(nums) <= 1 {
      return [][]int{nums}
    }
    
    ans := [][]int{}
        
    for i := 0; i < len(nums); i++ {
        v := make([]int, len(nums))    
        copy(v, nums)
      
        v[i], v[0] = v[0], v[i]
        v = v[1:]
    
        res := permute(v)
        
        for _, j := range res {
            j = append([]int{nums[i]}, j...)
            ans = append(ans, j)
        }
    }
    
    return ans
}
