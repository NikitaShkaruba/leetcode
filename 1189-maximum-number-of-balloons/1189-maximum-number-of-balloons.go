func maxNumberOfBalloons(text string) int {
  bb := "balloon"
  c := 0
  bChar := getCharCount(bb)
  t := getCharCount(text)

  for i := 0; i < len(text) / len(bb); i++ {
    for j := 0; j < 26; j++ {
      if bChar[j] > t[j] && bChar[j] != 0 {
        return c
      } else {
        t[j] = t[j] - bChar[j]
      }
    }
    
    c++
  }

  return c
}

func getCharCount(a string) []int {
  res := make([]int, 26)

  for _, v := range a {
    res[v - 'a']++
  }

  return res
}