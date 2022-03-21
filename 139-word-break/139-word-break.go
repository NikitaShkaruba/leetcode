/*

  Solutions:
    1. Backtrack: O(n^n) time, O(n) space
    2. Backtrack + trie: O(n^n) time, O(n) space, but better in terms of time on average
    3. Backtrack + trie + cahcing: O(n^n) time, O(n) space, but better in terms of time on average
    4. Backtrack + cache: O(n^n) time, O(n) space, but better in terms of time on average (implemented)
    
  Test cases:
    - s = "leetcode", wordDict = ["leet","code"] : True
    - s = "applepenapple", wordDict = ["apple","pen"] : True
    - s = "catsandog", wordDict = ["cats","dog","sand","and","cat"] : False
    - s = "a", wordDict = ["b"] : False
    - s = "a", wordDict = ["a"] : True (!!!)
    - s = "aaaaaaa", wordDict = ["a", "aa", "aaa", ...]
  
  Testing grounds:
    s = "applepenapple"
    
    wordDict = {
      "apple":1,
      "pen":1,
    }

*/
func wordBreak(s string, wordDict []string) bool {
  cache := make(map[string]struct{})  
  return backtrack(s, wordDict, cache)
}

func backtrack(s string, wordDict []string, cache map[string]struct{}) bool {
  if len(s) == 0 {
    return true
  }
  
  for _, w := range wordDict {
    if len(w) <= len(s) && s[:len(w)] == w {
      nextS := string(s[len(w):])
      
      if _, ok := cache[nextS]; ok {
        continue
      }
      
      if backtrack(nextS, wordDict, cache) {
        return true
      }
    }
  }
  
  cache[s] = struct{}{}
  
  return false
}