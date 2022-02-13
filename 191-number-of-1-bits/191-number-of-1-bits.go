/*

  Extra test cases:
    -

  Solutions:
    1. 
    
  Testing grounds:
    num: 00000000000000000000000000001011
    mask: 00000000000000000000000000000001
    
  
*/
func hammingWeight(num uint32) int {
  weight := 0
  var mask uint32 = 1
  
  for i := 0; i < 32; i++ {
    if num & mask > 0 {
      weight++
    }
    
    mask <<= 1
  }
  
  return weight
}