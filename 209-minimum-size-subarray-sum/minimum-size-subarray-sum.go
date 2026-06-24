func minSubArrayLen(target int, nums []int) int {
    left, sum := 0, 0
    minLength := math.MaxInt

    for right :=0; right <len(nums); right ++{
       sum = sum+ nums[right]
       
       for sum >= target {
        length := right-left+1
        if minLength > length{
            minLength = length
        }

        sum = sum-nums[left]       
        left ++
       }

    }
    if minLength  == math.MaxInt{
        return 0
    }
    return minLength
}

