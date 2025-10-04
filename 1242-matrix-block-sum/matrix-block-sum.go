func matrixBlockSum(mat [][]int, k int) [][]int {
    n := len(mat)
    m := len(mat[0])
    a := make([][]int, n)
    prefix_sum := make([][]int, n)

    for i :=range mat {
        a[i] = make([]int, m)
        prefix_sum[i] = make([]int, m)
    }

    // Префиксные суммы на двумерном массиве
    for i := range mat {
        for j := range mat[i] {
            prefix_sum[i][j] = mat[i][j]
            if i > 0 {
                prefix_sum[i][j] += prefix_sum[i-1][j]
            }
            if j > 0 {
                prefix_sum[i][j] += prefix_sum[i][j-1]
            }
            if i > 0 && j > 0 {
                prefix_sum[i][j] -= prefix_sum[i-1][j-1]
            }
        }
    }

    // Префиксные суммы на подматрице (двумерной)
    for i := range prefix_sum {
        for j := range prefix_sum[i] {
            a[i][j] = prefix_sum[min(n-1, i+k)][min(m-1, j+k)]
            if i - k > 0 {
                a[i][j] -= prefix_sum[max(0, i-k-1)][min(m-1, j+k)]
            }
            if j - k > 0 {
                a[i][j] -= prefix_sum[min(n-1, i+k)][max(0, j-k-1)]
            }
            if i - k > 0 && j - k > 0 {
                a[i][j] += prefix_sum[max(0, i-k-1)][max(0, j-k-1)]
            }
        }
    }
    return a
}