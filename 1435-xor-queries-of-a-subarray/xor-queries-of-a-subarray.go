func xorQueries(arr []int, queries [][]int) []int {
    n := len(arr)
    q := len(queries)
    prefix_sum := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix_sum[i+1] = prefix_sum[i] ^ arr[i]
    }

    ans := make([]int, q)
    for i := range queries {
        left, right := queries[i][0], queries[i][1]
        ans[i] = prefix_sum[right+1] ^ prefix_sum[left] // все, что идет до left - самоуничтожается
    }
    return ans
}