/**
  Testcases:
  - [flower, flowers, flowerd, flowere, gut]
  - [flower, flower, flower]
  - [flower, flow, flight] => fl
  - [dog, racecar, car] => ""
  - [dog, dog, ""] => ""
  - [dog] => dog (!!! +)
  - [dog, dog] => dog (!!! +)
  - ["", ""] => "" (!!! +)
  - [dog, ""] => "" (!!! +)
  
  Solutions:
  - Bruteforce: O(N*M) time, O(M) space, n is the amount of words, m is the max word length
  
    - Find a prefix between first two words
    - if only 2, return
    - else iterate over all the others, compact
      - if at the end, continue
      - if equal, iter++
      - if not equal, cut the prefix
        - if equal to zero, return
  
  Sandbox:
    [flower, flow, flight, flew, grew]
    flower, flow => flow
    flight flow => fl
      ^      ^
    flew fl
     ^    ^
    grew fl => ""
    ^    ^
    
  - [flower, flow, flight] => fl
*/
func longestCommonPrefix(strs []string) string {
  if strs == nil || len(strs) == 0 {
    return ""
  }
  if len(strs) == 1 {
    return strs[0]
  }
  
  commonPrefix := ""
  for i := 0; i < len(strs[0]) && i < len(strs[1]); i++ {
    if strs[0][i] != strs[1][i] {
      break
    }

    commonPrefix += string(strs[0][i])
  }
  
  if commonPrefix == "" {
    return commonPrefix
  }
  
  for i := 2; i < len(strs); i++ {
    curCommonPrefix := ""
    for j := 0; j < len(strs[i]) && j < len(commonPrefix); j++ {
      if strs[i][j] != commonPrefix[j] {
        break
      }

      curCommonPrefix += string(commonPrefix[j])
    }
    
    commonPrefix = curCommonPrefix
    if commonPrefix == "" {
      break
    }
  }
  
  return commonPrefix
}















