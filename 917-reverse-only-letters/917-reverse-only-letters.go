func reverseOnlyLetters(s string) string {
	c := []byte(s)
	i := 0
	j := len(s) - 1

	for i < j {
		for i < j && !isAlpha(s[i]) {
			i++
		}
		for j > i && !isAlpha(s[j]) {
			j--
		}
    
		c[i], c[j] = c[j], c[i]
    
		i++
		j--
	}

	return string(c)
}

func isAlpha(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}