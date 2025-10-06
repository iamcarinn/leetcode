func minStartValue(nums []int) int {
    n := len(nums)
    minSum := 0
    prefix_sum := make([]int, n+1)
    for i := range nums {
        prefix_sum[i + 1] = prefix_sum[i] + nums[i]
        if prefix_sum[i + 1] < minSum {
            minSum = prefix_sum[i+1]
        }
    }

    return 1 - minSum
    
}