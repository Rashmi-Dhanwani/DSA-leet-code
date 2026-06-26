func removeDuplicates(nums []int) int {
    duplicateMap := make(map[int]bool)
    j :=0

    for i , value := range nums {
       ok := duplicateMap[value]
        if ok {
           continue
        }
        duplicateMap[value]= true
        nums[j] = nums[i]
        j++
    }
    for i:=j;i<len(nums);i++{
        nums[i]= 0
    }
    return j
}


