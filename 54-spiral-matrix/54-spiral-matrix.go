/*

  Solutions:
    - Have separate "used" matrix: O(n) time, O(n) space
    - Set used numbers as 101: O(n) time, O(1) space

  Test cases:
    - [[ 1,  2,  3,  4]]
     - [[ 1,  2,  3,  4]
       [ 5,  6,  7,  8]]
     - [[ 1,  2,  3,  4]
       [ 5,  6,  7,  8]
       [ 9, 10, 11, 12]]
    - [[ 1,  2,  3,  4]
       [ 5,  6,  7,  8]
       [ 9, 10, 11, 12]
       [13, 14, 15, 16]]]
    - [[1, 2, 3]
       [4, 5, 6]
       [7, 8, 9]]]: [1, 2,3, 4, 5, 6, 7, 8, 9]
    - [[1]]
    - [[1],
       [2],
       [3]]
  
  Testing grounds:
    direction: 0
    - [[ 1000,  1000,  1000,  1000]
       [ 5,  6,  7,  8]
       [ 9, 10, 11, 12]]
       
*/

func spiralOrder(matrix [][]int) []int {
  matrixLen := len(matrix) * len(matrix[0])
  result := make([]int, 0, matrixLen)
  
  i := 0
  j := 0
  direction := 0 // 0 is right, 1 is down, 2 is left, 3 is up
  usedValue := 1000
  
  for len(result) != matrixLen {
    result = append(result, matrix[i][j])
    matrix[i][j] = usedValue
    
    nextI, nextJ := getNextIndex(i, j, direction)
    if nextI < 0 || nextI >= len(matrix) || nextJ < 0 || nextJ >= len(matrix[0]) || matrix[nextI][nextJ] == usedValue {
      direction = switchDirection(direction)
      nextI, nextJ = getNextIndex(i, j, direction)
    }
    
    i, j = nextI, nextJ
  }
  
  return result
}

func getNextIndex(i, j, direction int) (int, int) {
  switch direction {
  case 0:
    return i, j + 1
  case 1: 
    return i + 1, j
  case 2:
    return i, j - 1
  case 3:
    return i - 1, j
  default:
    return 0, 0
  }
}

func switchDirection(direction int) int {
  if direction == 3 {
    return 0
  }
  
  return direction + 1
}