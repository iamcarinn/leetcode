func runningSum(nums []int) []int {
    size := len(nums) + 1
    numsSum := make([]int, size)

    for i := range nums {
        numsSum[i+1] = numsSum[i] + nums[i]
    }
    return numsSum[1:]
}