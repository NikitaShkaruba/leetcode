/*

  Solutions:
    - Backtracking: O(2^2n) time, O(n) space (without the output space)
  
  Test cases:
    - bigwordbigwordbigwordbigword: ["15","14d"] (!!!)
    - word: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]
    - a: [1, a]
  
  Testing grounds:
    f(word, abbr)
    
    "", word
      "1", ord
        "1o", rd
          "1o1", d
          "1o2" (V)
          "1or", d
            "1ord" (V)
            "1or1" (V)
      "2", rd
        "2r", d
          "2rd" (V)
          "2r1",
      "3", d
        "3d" (V)
      "4" (V)
      "w", ord
        "wo", rd
          "wor", d
            "word" (V)
    
*/
func generateAbbreviations(word string) []string {
  res := []string{}
  backtrack(word, []byte{}, &res)
  return res
}

func backtrack(s string, abbr []byte, res *[]string) {
  if len(s) == 0 {
    *res = append(*res, string(abbr))
    return
  }
  
  abbreviateSomething := !(len(abbr) > 0 && isNumber(abbr[len(abbr)-1]))
  if abbreviateSomething {
    for i := 1; i <= len(s); i++ {
      backtrack(s[i:], append(abbr, intToChars(i)...), res)
    }  
  }
  
  backtrack(s[1:], append(abbr, s[0]), res)
}

func isNumber(b byte) bool {
  return b >= '0' && b <= '9'
}

func intToChars(i int) []byte {
  return []byte(strconv.Itoa(i))
}