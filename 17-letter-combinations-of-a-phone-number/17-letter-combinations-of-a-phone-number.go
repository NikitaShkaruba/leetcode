/*
  
  Solutions:
    - DFS: O(4^n) time, O(4^n) space
  
  Test cases:
    - "23": ["ad","ae","af","bd","be","bf","cd","ce","cf"]
    - "1234": [...]
    - "": [] (!!!)
    
  Testing ground:
    - 1234
      abcd
      
    - 23
    
                       "2345", ""
                    /       |        \
            "345", "a"   "345", "b   "345", "c
          /
    "45", "ad"    "45", "ae"
       /
    "5", "adg"
    /        \        \
  "", "adgj"  "", "adgk"  "", "adgl"
  
*/
func letterCombinations(digits string) []string {
  if len(digits) == 0 {
    return []string{}
  }
  
  return letterCombinationsHelper([]byte(digits), []byte{})
}

var digitLettersMap = map[byte]string{
  '2': "abc",
  '3': "def",
  '4': "ghi",
  '5': "jkl",
  '6': "mno",
  '7': "pqrs",
  '8': "tuv",
  '9': "wxyz",
}

func letterCombinationsHelper(digits []byte, combination []byte) []string {
  if len(digits) == 0 {
    return []string{string(combination)}
  }
  
  digitLetters := digitLettersMap[digits[0]]
  combinations := make([]string, 0)
  
  for i := 0; i < len(digitLetters); i++ {
    childCombinations := letterCombinationsHelper(digits[1:], append(combination, digitLetters[i]))
    combinations = append(combinations, childCombinations...)
  }
  
  return combinations
}













