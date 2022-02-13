func firstUniqChar(s string) int {
  if len(s) == 0 {
		return -1
	}

	charOccurencies := make([]int, 26)
	for _, v := range s {
		charOccurencies[byte(v)-'a']++
	}

	for i, v := range s {
		if charOccurencies[byte(v)-'a'] == 1 {
			return i
		}
	}
  
	return -1
}