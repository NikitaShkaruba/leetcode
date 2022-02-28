/*

  Solutions:
    - Brute force: O(a lot) time, O(a lot) space
    - DFS: O(n*4^l) time, O(l) space
  
  Test cases:
    - [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    - [[C]], C : true
    - [[C]], A : false
  
  Testing grounds:

*/

type IndexesPair struct {
  i int
  j int
}

func exist(board [][]byte, word string) bool {
  for i := range board {
    for j := range board[i] {
      if board[i][j] != word[0] {
        continue
      }
      
      seen := make(map[IndexesPair]struct{})
      if dfs(i, j, board, word, seen) {
        return true
      }
    }
  }
  
  return false
}

func dfs(i int, j int, board [][]byte, word string, seen map[IndexesPair]struct{}) bool {
  if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
    return false
  }

  if board[i][j] != word[0] {
    return false
  }
  
  ijPair := IndexesPair{i, j}
  if _, exists := seen[ijPair]; exists {
    return false
  }
  seen[ijPair] = struct{}{}
  
  if len(word) - 1 == 0 {
    return true
  }
  
  wasFound := dfs(i-1, j, board, word[1:], seen) ||
              dfs(i+1, j, board, word[1:], seen) ||
              dfs(i, j+1, board, word[1:], seen) ||
              dfs(i, j-1, board, word[1:], seen)
  
  delete(seen, ijPair)
  
  return wasFound
}