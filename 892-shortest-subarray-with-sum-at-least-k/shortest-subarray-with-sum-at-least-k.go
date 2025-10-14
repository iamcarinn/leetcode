func shortestSubarray(nums []int, k int) int {
    n := len(nums)

    prefix_sum := make([]int, n + 1)
    for i := range nums {
        prefix_sum[i+1] = prefix_sum[i] + nums[i]
    }

    result := n + 1
    deque := []int{}

    for i := 0; i <= n; i++ {
        // Удаляем из начала deque, пока сумма >= k
        for len(deque) > 0 && prefix_sum[i]-prefix_sum[deque[0]] >= k {
            if i - deque[0] < result {
                result = i - deque[0]
            }
            deque = deque[1:]
        }

        // Удаляем из конца deque все индексы, у которых prefix_sum больше или равен текущего
        for len(deque) > 0 && prefix_sum[i] <= prefix_sum[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }

        deque = append(deque, i)
    }

    if result == n+1 {
        return -1
    }
    return result

}