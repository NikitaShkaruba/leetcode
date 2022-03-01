/*

  Solutions:
    - DFS: O(n*2^n) time, O(n) space

  Test cases:
    - a1b2: ["a1b2","a1B2","A1b2","A1B2"]
    - 1: [1]
    - a: [a, A]
  
  Testing grounds:
  
    a1b2
    ^
  a, a1b2    A, a1b2
      ^          ^
  a1, a1b2    A1, a1b2
        ^           ^
 a1b a1B         A1b  A1B
 a1b2 a1B2       A1b2  A1B2
*/

func letterCasePermutation(s string) []string {
  permutations := []string{}
  dfs(s, 0, []byte{}, &permutations)
  return permutations
}

func dfs(s string, i int, perm []byte, permutations *[]string) {
  if len(s) == i {
    *permutations = append(*permutations, string(perm))
    return
  }
  
  if isNumber(s[i]) {
    dfs(s, i+1, append(perm, s[i]), permutations)
  } else {
    dfs(s, i+1, append(perm, s[i]), permutations)
    dfs(s, i+1, append(perm, switchCase(s[i])), permutations)
  }
}

func isNumber(b byte) bool {
  return b >= '0' && b <= '9'
}

func switchCase(b byte) byte {
  if b >= 'a' && b <= 'z' {
    return 'A' + (b - 'a')
  } else {
    return 'a' + (b - 'A')
  }
}