func rotate(matrix [][]int)  {
  n := len(matrix) - 1
  
  for x := 0; x < len(matrix)/2; x++ {
    for y := x; y < n-x; y++ {
      t := matrix[x][y]
      
      matrix[x][y] = matrix[n-y][x]
      matrix[n-y][x] = matrix[n-x][n-y]
      matrix[n-x][n-y] = matrix[y][n-x]
      matrix[y][n-x] = t
    }
  }
}