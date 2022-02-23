/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

    Solutions:
        1. DFS: O(n) time, O(n) space
    
    Test cases:
        - [1]: 0 (!!!)
        - [1, 2, 3]: 2
        - [1, 2, 3, 4, 5]: 3
    
    Testing grounds:

*/
func diameterOfBinaryTree(root *TreeNode) int {
    diameter, _ := dfs(root)
    return diameter
}

func dfs(node *TreeNode) (int, int) {
    ourDiameter := 0
    
    leftDiameter, leftMaxPath := 0, 0
    if node.Left != nil {
        leftDiameter, leftMaxPath = dfs(node.Left)
        leftMaxPath++
        ourDiameter += leftMaxPath
    }
    
    rightDiameter, rightMaxPath := 0, 0
    if node.Right != nil {
        rightDiameter, rightMaxPath = dfs(node.Right)    
        rightMaxPath++
        ourDiameter += rightMaxPath
    }
    
    maxDiameter := ourDiameter
    if leftDiameter >= rightDiameter && leftDiameter >= ourDiameter {
        maxDiameter = leftDiameter
    } else if rightDiameter >= leftDiameter && rightDiameter >= ourDiameter {
        maxDiameter = rightDiameter
    }
    
    maxPath := leftMaxPath
    if rightMaxPath > leftMaxPath {
        maxPath = rightMaxPath
    }
    
    return maxDiameter, maxPath
}