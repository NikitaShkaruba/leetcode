/*
  
  Solutions:
    - Brute force: O(n^2*m*l) time, O(1) space
    - Trie: O(m*l + n^2*l) time, O(m*l) space
    
    [n is the len of text, m is the len of words, l is the len of the biggest word]
  
  Test cases:
    - text = "thestoryofleetcodeandme", words = ["story","fleet","leetcode"] : [[3,7],[9,13],[10,17]]
    - text = "ababa", words = ["aba","ab"] : [[0,1],[0,2],[2,3],[2,4]]
    - text = "a", words = ["a""] : [[0,0]]
  
  Testing grounds:
    "ababa"

      *
    a  b
    b*  a 
    a  b*
    b* a
    a*  *
    *
    
    ababa, [aba, ab]
    
    a 
    b 
    a* 
    *

*/

type TrieNode struct {
  children map[byte]*TrieNode
}

type Trie struct {
  root *TrieNode
}

func (t *Trie) insert(word string) {
  node := t.root
  
  for _, b := range []byte(word) {
    if _, contains := node.children[b]; !contains {
      node.children[b] = &TrieNode{ children: make(map[byte]*TrieNode) }
    }
    
    node = node.children[b]
  }
  
  // Mark the end of the word
  if _, contains := node.children['*']; !contains {
    node.children['*'] = &TrieNode{ children: make(map[byte]*TrieNode) }
  }
}

func (t *Trie) find(word string) bool {
  node := t.root
  
  for _, b := range []byte(word) {
    if _, contains := node.children[b]; !contains {
      return false
    }
    
    node = node.children[b]
  }
  
  // We we want only fully included words, not their parts
  if _, contains := node.children['*']; contains {
    return true
  }
  
  return false
}

func indexPairs(text string, words []string) [][]int {
  trie := Trie{ root: &TrieNode{ children: make(map[byte]*TrieNode) } }
  for _, w := range words {
    trie.insert(w)
  }
  
  pairs := [][]int{}
  for i := 0; i < len(text); i++ {
    for j := i + 1; j <= len(text); j++ {
      if trie.find(text[i:j]) {
        pairs = append(pairs, []int{i, j-1})
      }      
    }
  }
  
  return pairs
}
