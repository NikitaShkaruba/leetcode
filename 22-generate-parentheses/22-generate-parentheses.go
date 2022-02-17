/*

  Solution:
    1. DFS: O(n * 2^n) time, O(2^n * n) space
  
  Test cases:
    - 1: ["()"]
    - 3: ["((()))","(()())","(())()","()(())","()()()"]
  
  Testing grounds:
            
                        f(3, 0, "")
                      /                          \
                f(2, 1, "(")                       f(2, 0, "()")
                /             \                          /    \
      f(1, 2, "((")            f(1, 1, "(()")       1, 1      1, 0
    /             \                |     \         /    \    /    \  
  !!!!      f(0, 2, "((()")   .      .       .     .    .      . 
  
  f(1, 0, "")
*/
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	result = genHelper(n, 0, []byte{}, result)
	return result
}

func genHelper(toOpen int, toClose int, curSequence []byte, result []string) []string {
	if toOpen == 0 && toClose == 0 {
		result = append(result, string(curSequence))
		return result
	}

	if toOpen > 0 {
		result = genHelper(toOpen-1, toClose+1, append(curSequence, '('), result)
	}

	if toClose > 0 {
		result = genHelper(toOpen, toClose-1, append(curSequence, ')'), result)
	}

	return result
}

