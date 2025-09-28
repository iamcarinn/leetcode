func isCovered(ranges [][]int, left int, right int) bool {
    // проставляю 1 и -1 для старт и енд
    // считаю префиксную сумму
    // если хотя бы один 0 на отрезке в префиксной сумме -> false

    oneArr := make([]int, 52)
    for i := range ranges {
        start, end := ranges[i][0], ranges[i][1]
        oneArr[start] += 1
        oneArr[end+1] -= 1
    }

    prefixSum := make([]int, len(oneArr)+1)
    for i := range oneArr {
        prefixSum[i+1] = prefixSum[i] + oneArr[i]
    }

    for i := left; i <= right; i++ {
        if  prefixSum[i+1] == 0 {
            return false
        }
    }

    return true
}