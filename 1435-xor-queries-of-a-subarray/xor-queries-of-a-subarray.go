func xorQueries(arr []int, queries [][]int) []int {
    n := len(arr)
    q := len(queries)
    prefix_xor := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix_xor[i+1] = prefix_xor[i] ^ arr[i]
    }

    ans := make([]int, q)
    for i := range queries {
        left, right := queries[i][0], queries[i][1]
        ans[i] = prefix_xor[right+1] ^ prefix_xor[left] // все, что идет до left - самоуничтожается
    }
    return ans
}