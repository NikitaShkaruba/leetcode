func merge(nums1 []int, m int, nums2 []int, n int) {
  i := m-1
  j := n-1
	r := m + n - 1

	for r >= 0 {
		if i >= 0 && (j < 0 || nums1[i] >= nums2[j]) {
			nums1[r] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[r] = nums2[j]
			j--
		}
    
		r--
	}
}