//        1

//       1, 1
//      [0, 1]

//     1, 2, 1
//    [0, 1, 2] => Math.Ceil(3 / 2)

//    1, 3, 3, 1
//  [ 0, 1, 2, 3] => Math.Ceiling(4 / 2) = 2

//   1, 4, 6, 4, 1
//  [0, 1, 2, 3, 4] => Math.Ceiling(5 / 2) = 3

//  1, 5, 10,10, 5, 1
// [0, 1, 2, 3,  4, 5] 6 - i - 1 =   

// 1, 6, 15, 20, 15, 6, 1
//[0, 1, 2,  3,  4,  5, 6]
func generate(numRows int) [][]int {
  triangle := make([][]int, 0, numRows)
  triangle = append(triangle, []int{1})
  
  for i := 1; i < numRows; i++ {
    tierLen := (len(triangle[i-1]) + 1)
    
    triangle = append(triangle, make([]int, tierLen))
    
    for j := 0; j < tierLen; j++ {
      if (j == 0 || j == tierLen - 1) {
        triangle[i][j] = 1  
      } else {
        triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
      }
    }
  }
  
  return triangle
}