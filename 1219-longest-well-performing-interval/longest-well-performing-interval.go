func longestWPI(hours []int) int {
    // интервал, где переработок больше недороботок
    need_to_work := 8
    n := len(hours)

    table := make([]int, n)
    for i := range hours {
        if hours[i] > need_to_work {
            table[i] = 1
        } else {
            table[i] = -1
        }
    }
    // [1 1 -1 -1 -1 -1 1]

    prefix_sum := make([]int, n+1)
    for i := range table {
        prefix_sum[i+1] = prefix_sum[i] + table[i]
    }
    // [0 1 2 1 0 -1 -2 -1]

    max_len := 0
    for i := range prefix_sum {
        for j := n; j >= i + 1; j-- {
            if prefix_sum[j]-prefix_sum[i] > 0 {
                max_len = max(max_len, j-i)
            }
        }
    }   
    return max_len
}