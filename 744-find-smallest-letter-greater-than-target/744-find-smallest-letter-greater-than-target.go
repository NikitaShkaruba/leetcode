/*

  Solutions:
    1. Brute force: O(n) time, O(1) space
    2. Binary search: Avg: {O(logn) time, O(1) space}, Worst: {O(n) time, O(1) space}
  
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

  // The only way when we return letters[0] is when either we:
  // - we were hitting equals at the end all the time and got to the end - we need to wrap. E.g. {[1, 2, 3, 4, 5, 5, 5], 5}
  // - we were hitting less   at the end all the time and got to the end - we need to wrap. E.g. {[1, 2, 3, 4, 5, 6, 7], 8}
  //
  // If we've gotten at least once to the `right = middle` result, we'll always find the greater target.
  if letters[left] <= target {
    return letters[0]
  }
  
  return letters[left]  
}