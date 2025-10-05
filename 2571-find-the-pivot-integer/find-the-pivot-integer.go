func pivotInteger(n int) int {
    arr := make([]int, n)
    for i := range arr {
        arr[i] = i + 1
    }

    prefix_sum := make([]int, n + 1)
    for i := range arr {
        prefix_sum[i+1] = prefix_sum[i] + arr[i]
    }
    ans := -1
    for x := 1; x <= n; x++ {
        if prefix_sum[x] == prefix_sum[n] - prefix_sum[x-1] {
            ans = x
        }
    }
    return ans
}