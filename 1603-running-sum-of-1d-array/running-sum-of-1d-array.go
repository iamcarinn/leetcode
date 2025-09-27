func runningSum(nums []int) []int {
    size := len(nums)
    numsSum := make([]int, size)
    numsSum[0] = nums[0]
    
    for i := 0; i < size - 1; i++ {
        numsSum[i+1] = numsSum[i] + nums[i+1]
    }
    return numsSum
}