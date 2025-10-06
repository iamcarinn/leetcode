func answerQueries(nums []int, queries []int) []int {
    n := len(nums)
    m := len(queries)
    answer := make([]int, m)

    // Сортируем nums, чтобы одти от меньших
    sort.Ints(nums)

    // Находим префиксные суммы
    prefix_sum := make([]int, n + 1)
    prefix_sum[0] = 0
    for i := range nums {
        prefix_sum[i+1] = prefix_sum[i] + nums[i]
    }

    // Для каждого queries[i]
    for i := range queries {
        lenOK := 0
        for lenOK+1 <= n && prefix_sum[lenOK+1] <= queries[i] {
            lenOK++
        }
        answer[i] = lenOK
    }

    return answer
}