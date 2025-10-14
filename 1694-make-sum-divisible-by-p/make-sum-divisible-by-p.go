func minSubarray(nums []int, p int) int {
    n := len(nums)
    prefix_sum := make([]int, n + 1)
    for i := range nums {
        prefix_sum[i+1] = prefix_sum[i] + nums[i]
    }

    if prefix_sum[n] % p == 0 {
        return 0
    }
    if prefix_sum[n] < p {
        return -1
    }

    for i := range prefix_sum {
        for j := 0; j <= n - i; j++ {
            if i != n && (prefix_sum[n] - prefix_sum[i+j] + prefix_sum[j]) % p == 0 {
                return i
            }
        }
    }

    return -1
}