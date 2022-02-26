/*
  
  Solutions:
    - sort them by start date: O(n*logn) time, O(logn) space
  
  Test cases:
    - [[0,5],[0,30]]: false
    - [[0,5],[6,10]]: true
    - [[0,5],[5,10]]: true
    - [[0,30],[5,10],[15,20]]: false
    - [[7,10],[2,4]]: true
    - [0, 1][0, 2]: false
    - []: true (!!!)
  
  Testing grounds:
  
*/

func canAttendMeetings(intervals [][]int) bool {
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  
  for i := 0; i < len(intervals) - 1; i++ {
    if intervals[i][1] > intervals[i+1][0] {
      return false
    }
  }
  
  return true
}