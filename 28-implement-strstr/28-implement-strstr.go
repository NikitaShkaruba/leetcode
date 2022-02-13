/**
  * -1 if not a part
  * 0 if empty string
  
  Test cases:
    5, 2
  * hello, ll => 2
  * aaaaa, bba => -1
  * "", "" => 0
  * llllllo, llllllb => 2
  * aa, bbaaaaa => -1 (!!! +)
  * aaaa, "" => -1 (!!! +)
  
  Solutions:
    * Brure force: O(nm) time, O(1) space, n = first string len, m = second string len
    
  Sandbox:
    hello, ll
      i     j
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

  outerloop:
  for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		for k := 0; k < len(needle); k++ {
			if needle[k] != haystack[i+k] {
				continue outerloop
			}
		}

		return i
	}

	return -1
}