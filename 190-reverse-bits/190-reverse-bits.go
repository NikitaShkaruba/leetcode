/**

  Test cases
    - 000000010100101000001111010011100 -> 00111001011110000010100101000000

  Testing ground:
  
    1 iteration
      num: 000000010100101000001111010011101
      mask: 000000000000000000000000000000001
      reversedMask: 100000000000000000000000000000000
      reversedNum: 100000000000000000000000000000000
      
    2 iteration
      num: 000000010100101000001111010011101
      mask: 000000000000000000000000000000010
      reversedMask: 010000000000000000000000000000000
      reversedNum: 100000000000000000000000000000000
      
    3 iteration
      num: 000000010100101000001111010011101
      mask: 000000000000000000000000000000100
      reversedMask: 001000000000000000000000000000000
      reversedNum: 101000000000000000000000000000000
    
  Solutions:
    - Brute force: hard
    - Two masks: O(n) time, O(1) space
*/
func reverseBits(num uint32) uint32 {
  var reversedNum uint32
  
  var mask uint32 = 1 
  var reversedMask uint32 = 1 << 31
  
  for i := 0; i < 32; i++ {
    if num & mask > 0 {
      reversedNum |= reversedMask
    }
    
    mask <<= 1
    reversedMask >>= 1
  }
  
  return reversedNum
}