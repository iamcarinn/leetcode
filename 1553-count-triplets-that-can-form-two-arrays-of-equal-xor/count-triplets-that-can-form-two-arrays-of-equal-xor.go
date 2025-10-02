func countTriplets(arr []int) int {
    n := len(arr)
    prefix_xor := make([]int, n)
    prefix_xor[0] = arr[0]

    // матрица префиксных xor
    for i := 1; i < n; i++ {
        prefix_xor[i] = prefix_xor[i-1] ^ arr[i]
    }

    ans := 0
    left_xor := 0
    right_xor := 0
    for i := range arr {
        for j := i + 1; j <= n; j++ {
            for k := j; k < n; k++ {
                left_xor = prefix_xor[j - 1]
                if i - 1 >= 0 {
                    left_xor ^= prefix_xor[i - 1]
                }
                right_xor = prefix_xor[k] ^ prefix_xor[j - 1]
                if left_xor == right_xor {
                    ans += 1
                }
            }
        }
    }
    return ans
}