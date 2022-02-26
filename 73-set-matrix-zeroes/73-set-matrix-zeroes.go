/*

  Solutions:
    - Mark all values to modify: O(n*m + (n*m)^2) time, O(n*m) space
    - Store rows and columns to modify in set: O(n*m + (n*m)^2) time, O(n+m) space
    - Store rows and columns to modify in initial matrix: O(n*m) time, O(1) space
  
  Test cases:
    - [[1,1,1],[1,0,1],[1,1,1]] : [[1,0,1],[0,0,0],[1,0,1]] (One zero at the middle)
    - [[0,1,2,0],[3,4,5,2],[1,3,1,5]] : [[0,0,0,0],[0,4,5,0],[0,3,1,0]] (Two intercepted rows)
    - [[0,0,0],[0,0,0],[0,0,0]] : [[0,0,0],[0,0,0],[0,0,0]] (All zeroes)
    - [[1]] : [[1]] (!!!)
    - [[0]] : [[0]] (!!!)
  
  Testing grounds:
    0, 1, 0, 1, 0
    1, 1, 1, 1, 0
    1, 1, 1, 1, 0
    0, 1, 0, 1, 0
    1, 1, 1, 1, 1
    
    1,0,1
    0,0,0
    1,1,1
    
    1,   2,  3,  4
    5,   0,  7,  8
    0,  10, 11, 12
    13, 14, 15,  0
*/

func setZeroes(matrix [][]int) {
  // Understand
  firstColumnNeedsZeroing := false
  for i := 0; i < len(matrix) && !firstColumnNeedsZeroing; i++ {
    if matrix[i][0] == 0 {
      firstColumnNeedsZeroing = true
    }
  }
  
  firstRowNeedsZeroing := false
  for j := 0; j < len(matrix[0]) && !firstRowNeedsZeroing; j++ {
    if matrix[0][j] == 0 {
      firstRowNeedsZeroing = true
    }
  }
  
  // Store facts that the row and column needs to be zeroes if the first row and column
  for i := 1; i < len(matrix); i++ {
    for j := 1; j < len(matrix[0]); j++ {
      if matrix[i][j] == 0 {
        matrix[i][0] = 0
        matrix[0][j] = 0
      }
    }
  }

  // Iterate over stored facts of the rows, and set them to zeroes
  for i := 1; i < len(matrix); i++ {
    if matrix[i][0] == 0 {
      for j := 1; j < len(matrix[0]); j++ {
        matrix[i][j] = 0
      }
    }
  }
  
  // Iterate over stored facts of the columns, and set them to zeroes
  for j := 1; j < len(matrix[0]); j++ {
    if matrix[0][j] == 0 {
      for i := 1; i < len(matrix); i++ {
        matrix[i][j] = 0
      }
    }
  }
  
  // Don't forget to set the first column and row to zeroes if needed
  if firstColumnNeedsZeroing {  
    for i := 0; i < len(matrix); i++ {
      matrix[i][0] = 0
    }
  }
  if firstRowNeedsZeroing {
    for j := 0; j < len(matrix[0]); j++ {
      matrix[0][j] = 0
    }
  }
}