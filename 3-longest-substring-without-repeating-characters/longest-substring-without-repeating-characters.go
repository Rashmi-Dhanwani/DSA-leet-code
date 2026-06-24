func lengthOfLongestSubstring(s string) int {

    left :=0
    minLength := 0
    seen := make(map[byte]bool)

    for right := 0; right< len(s); right ++{
      for seen[s[right]]{
        delete(seen, s[left])
        left++
      }
      seen[s[right]]=true
      if right -left +1 > minLength{
        minLength = right -left +1
      }
       }
       return minLength
    
}


