/*

  Solutions:
    - Brute force: O(a lot) time, O(a lot) space
    - Backtracking: O(n*3^l) time, O(l) space. If you want to have lock-free algorithm, use a separate seen hashset
    - Backtracking + Pruning (caching the incorrect paths): O(n*3^l) time, O(l + n) space. This might perform better on a very large boards and words
  
  Test cases:
    - [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    - [[C]], C : true
    - [[C]], A : false
  
  Testing grounds:

*/

func exist(board [][]byte, word string) bool {
  for i := range board {
    for j := range board[i] {
      if board[i][j] != word[0] {
        continue
      }
      
      if backtrack(i, j, board, word) {
        return true
      }
    }
  }
  
  return false
}

func backtrack(i int, j int, board [][]byte, word string) bool {
  if len(word) == 0 {
    return true
  }
  
  if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
    return false
  }

  if board[i][j] != word[0] {
    return false
  }
  
  board[i][j] = '#'
  
  wasFound := backtrack(i-1, j, board, word[1:]) ||
              backtrack(i+1, j, board, word[1:]) ||
              backtrack(i, j+1, board, word[1:]) ||
              backtrack(i, j-1, board, word[1:])
  
  board[i][j] = word[0]
  
  return wasFound
}