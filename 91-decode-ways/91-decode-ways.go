/*
  
  Solutions:
    - Check for leading zeroes + backtracking + cache (Top down): O(n) time, O(n) space
    - Check for leading zero + dynamic programming (Bottom up): O(n) time, O(n^2) space
  
  Test cases:
    - "12": 2
    - "226": 3
    - "06": 0 (!!!)
    - "006": 0 (!!!)
    - "100005": 0 (!!!)
  
  Testing grounds:
    226
    /  \
   22  2
  /  \  \
 2    *  *
 \
  *
  
   106
   /
  10
  \
   *
   
   006
   / \
  00
  /
  
  22 6 - 226 - (3)
  2 2 6 - 226
  2 26 - 226
  
  f(string, map[string]int): int
  
  "2263"
  ""
  "2"
  "22"
  "226": "22", "2"
  "2263": "226"
  ["":1,"2":1,"22":2,"226":3,"2263":3]
  
  "100005"
  "": 1
  "1": "" : 1
  "10": "" : 1
  "100": : 0
*/

var decodeSet map[string]struct{} = map[string]struct{} {
  "1": struct{}{},
  "2": struct{}{},
  "3": struct{}{},
  "4": struct{}{},
  "5": struct{}{},
  "6": struct{}{},
  "7": struct{}{},
  "8": struct{}{},
  "9": struct{}{},
  "10": struct{}{},
  "11": struct{}{},
  "12": struct{}{},
  "13": struct{}{},
  "14": struct{}{},
  "15": struct{}{},
  "16": struct{}{},
  "17": struct{}{},
  "18": struct{}{},
  "19": struct{}{},
  "20": struct{}{},
  "21": struct{}{},
  "22": struct{}{},
  "23": struct{}{},
  "24": struct{}{},
  "25": struct{}{},
  "26": struct{}{},
}

func numDecodings(s string) int {
  dp := make([]int, len(s)+1)
  dp[0] = 1
  
  for i := 1; i <= len(s); i++ {
    subS := s[0:i]
  
    sum := 0

    cutIdx := len(subS)-2
    if cutIdx >= 0 {
      cut := subS[cutIdx:]
      if _, ok := decodeSet[cut]; ok {
        dpIdx := len(subS[0:cutIdx])
        sum += dp[dpIdx] 
      }
    }
    
    cutIdx = len(subS)-1
    if cutIdx >= 0 {
      cut := subS[cutIdx:]
      if _, ok := decodeSet[cut]; ok {
        dpIdx := len(subS[0:cutIdx])
        sum += dp[dpIdx]
      }
    }
    
    // zeroes at the beginning
    if sum == 0 {
      return 0
    }
    
    dp[len(subS)] = sum
  }
  
  return dp[len(s)]
}


















