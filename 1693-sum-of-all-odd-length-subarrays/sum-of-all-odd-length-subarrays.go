func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    ans := 0
    prefix_sum := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix_sum[i+1] = prefix_sum[i] + arr[i]
    }


    for lenInterval := 1; lenInterval <= n; lenInterval += 2 {
        for i := 0; i+lenInterval <= n; i++ {
            ans += prefix_sum[i+lenInterval] - prefix_sum[i]
        }
    }

    return ans
}