func isCovered(ranges [][]int, left int, right int) bool {
    // Размер массива должен быть достаточным для значений, которые можно ожидать
    oneArr := make([]int, 102) // берём больше, чтобы гарантировать хватит места для индексов

    for _, r := range ranges {
        start, end := r[0], r[1]
        oneArr[start]++
        if end+1 < len(oneArr) {
            oneArr[end+1]--
        }
    }

    prefixSum := make([]int, len(oneArr))
    for i := 1; i < len(oneArr); i++ {
        prefixSum[i] = prefixSum[i-1] + oneArr[i]
    }

    // Проверяем все значения на отрезке от left до right
    for i := left; i <= right; i++ {
        if prefixSum[i] == 0 {
            return false
        }
    }

    return true
}