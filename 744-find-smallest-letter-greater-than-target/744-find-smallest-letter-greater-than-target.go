/*

  Solutions:
    1. Brute force: O(n) time, O(1) space
    2. Binary search: O(logn) time, O(1) space
  
  Test cases:
    letters = ["c","f","j"], target = "a" : c
    letters = ["c","f","j"], target = "z" : 
    letters = ["c","f","j"], target = "c" : f
    letters = ["c","f","j"], target = "d" : f
    letters = ["a","b"], target = "z" : a

  Testing grounds:
    ["c","f","h"], j
    ["c","f","j"], a
    ["c","f","j"], f
    ["a","b"]
            ^
    min: b
    
    a, b, c, d, e, f, g, h, j*, k, l  ==== i
                      r   lm
                         
    letters = ["c","f","j"], target = "a" : c
                lr
    letters = ["c","f","j"], target = "z" : 
                        lr
    letters = ["c","f","j"], target = "c" : f
                lr
    letters = ["c","f","j"], target = "d" : f
                lr       
               
    possibleMin: j
  
*/
func nextGreatestLetter(letters []byte, target byte) byte {
  left := 0
  right := len(letters) - 1
 
  for left < right {
    middle := left + (right - left) / 2
    
    if letters[middle] <= target {
      left = middle + 1
    } else {
      right = middle
    }
  }

  if letters[left] > target {
    return letters[left]
  } else {
    return letters[0]
  }
}