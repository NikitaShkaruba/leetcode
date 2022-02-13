/**
  Notes:
    - Each prerequisite [a, b] says that you need to finish b before a
    
  Test cases:
  
  1.
              7 <- 9
              ^    v
    1 -> 4 -> 5 -> 6 
    
  2. 
              7
              ^
    1 -> 4 -> 5 <-> 6
    
  3. empty prerequisites => return true
  
  4. Loops?
  
  Solutions:
    1. DFS + counting looped courses: O(v + e) time, O(v + e) space. Not working.
    2. DFS + tracking current visiting cources: O(v + e) time, O(v + e) space
    
  Plan:
    * Initialize prerequisitesMap
    * Start from the cource 0, go through edges until you hit the base case with DFS.
      * For every edge remember it in current set
      * If we'd confronted an already visited node, return false
    * Return true
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesMap := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
    prerequisitesMap[i] = make([]int, 0)
	}
	for _, p := range prerequisites {
		from := p[0]
		to := p[1]
		prerequisitesMap[from] = append(prerequisitesMap[from], to)
	}

  seenSet := make(map[int]struct{})
  
  for courseId := 0; courseId < numCourses; courseId++ {
    canFinish := depthFirstSearch(courseId, prerequisitesMap, seenSet)  
    if !canFinish {
      return false
    }
  }
  
  return true
}

func depthFirstSearch(courseId int, prerequisitesMap map[int][]int, seenSet map[int]struct{}) bool {
  if _, seen := seenSet[courseId]; seen {
    return false
  }
  
  seenSet[courseId] = struct{}{}
  defer delete(seenSet, courseId)
  
  coursePrerequisites := prerequisitesMap[courseId]
  
  if len(coursePrerequisites) == 0 {
    return true
  }
  
  for _, prereqCourseId := range coursePrerequisites {
    canFinish := depthFirstSearch(prereqCourseId, prerequisitesMap, seenSet)
    if !canFinish {
      return false
    }
  }
  
  prerequisitesMap[courseId] = prerequisitesMap[courseId][:0] 
  
  return true
}


