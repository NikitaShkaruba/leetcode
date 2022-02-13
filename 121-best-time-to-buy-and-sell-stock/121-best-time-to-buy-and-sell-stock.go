/**
  Notes:
    - return 0 if no profit can be achieved
  
  Test cases:
    - [7, 1, 5, 3, 6, 4] => 5
    - [7, 6, 5, 3, 1, 4] => 5
    - [7, 6, 4, 3, 1] => 0
    - [7, 6, 4, 3, 3, 1] => 0
    - [7] => 0 (!!!)
    
  Testing grounds:
    [7, 1, 5, 3, 6, 4]
    [0, 5, 1, 3, 2, 0]
    
    [7, 1, 5, 3, 6, 4]
    [ ,  ,  ,  ,  , 0]
    
    [7, 1, 5, 3, 6, 4] => 5
    [7, 1, 5, 3, 6, 4]
    
    #               #
    #         #     #
    #   #     #     #
    #   #     # #   #
    #   # # # # #   #
    #   # # # # #   #
    # # # # # # #   #
    7 1 5 3 3 6 4 0 7
              ^
    
    min: 1
    max: 6
    maxProfit: 5
    
    #
    # #      
    # #    
    # # #   
    # # # # 
    # # # # 
    # # # # #
    7 6 4 3 1
    
    min: 1
    max: -1
    maxProfit: 0
    
  Solutions:
    - Brute force: O(n^2) time, O(1) space
    - One iteration: O(n) time, O(1) space
*/
func maxProfit(prices []int) int {
  min := 0
  max := -1
  maxProfit := 0
  
  for i := 1; i < len(prices); i++ {
    if prices[i] == prices[min] {
      continue
    }
    
    if prices[i] < prices[min] {
      min = i
      max = -1
    } else {
      if max == -1 || prices[i] > prices[max] {
        max = i
        
        newProfit := prices[max] - prices[min]
        if newProfit > maxProfit {
          maxProfit = newProfit
        }
      }
    }
  }
  
  return maxProfit
}

















