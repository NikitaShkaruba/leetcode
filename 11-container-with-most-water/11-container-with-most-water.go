/*

  Test cases:
    - [1, 8, 6, 2, 5, 4, 8, 3, 7]: 49
    - [1, 2, 3]: 1
    - [3, 2, 1]: 1
    - [1, 1]: 3
    - [0, 0]
     
  Solutions:
    - Brute force: O(n^2) time, O(1) space
    - Two pointers: O(n) time, O(1) space
  
  Testing grounds:
    - [1, 8, 6, 2, 5, 4, 8, 3, 7]
          l                    r
          1                    8
*/
func maxArea(height []int) int {
  l := 0
  r := len(height) - 1
  maxArea := 0
  
  for l < r {
    area := calculateArea(l, r, height)
    
    if area > maxArea {
      maxArea = area
    }
    
    if height[l] < height[r] {
      l++
    } else {
      r--
    }
  }
  
  return maxArea
}

func calculateArea(l, r int, height []int) int {
  var min int
  if height[l] < height[r] {
    min = height[l]
  } else {
    min = height[r]
  }
  
  return min * (r - l)
}