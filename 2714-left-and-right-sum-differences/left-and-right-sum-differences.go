func leftRightDifference(nums []int) []int {
    size := len(nums)
    leftSum := make([]int, size)
    rightSum := make([]int, size)
    result := make([]int, size)

    for i := 1; i < size; i++ {
        leftSum[i] = leftSum[i-1] + nums[i-1]
    }

    for j := size - 2; j >= 0; j-- {
        rightSum[j] = rightSum[j+1] + nums[j+1]
    }

    for k := range result {
        result[k] = leftSum[k] - rightSum[k]
        if result[k] < 0 {
            result[k] = -result[k]
        }
    }
    return result
}