func answerQueries(nums []int, queries []int) []int {
    n := len(nums)
    m := len(queries)
    answer := make([]int, m)

    // Сортируем nums
    sort.Ints(nums)

    // Считаем префиксные суммы
    prefix_sum := make([]int, n + 1)
    prefix_sum[0] = 0
    for i := range nums {
        prefix_sum[i+1] = prefix_sum[i] + nums[i]
    }
    // [ 1 2 4 5]
    // [ 0 1 3 7 12]




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