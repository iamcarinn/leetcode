func findMiddleIndex(nums []int) int {
    prefixSum := make([]int, len(nums)+1)
    suffixSum := make([]int, len(nums)+1)

    for i := range nums {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }

    for i := len(nums) - 1; i >= 0; i-- {
        suffixSum[i] = suffixSum[i+1] + nums[i]
    }

    for i := range nums {
        left_sum := prefixSum[i]       // от 0 до i-1
        right_sum := suffixSum[i+1]    // от i+1 до конца

        if left_sum == right_sum {
            return i
        }
    }

    return -1
}
