func numberOfPoints(nums [][]int) int {
    cars := make([]int, 102)

    for i := range nums {
        start, end := nums[i][0], nums[i][1]
        cars[start] += 1
        cars[end+1] -= 1
    }

    prefixSum := make ([]int, len(cars)+1)
    prefixSum[0] = 0
    counter := 0
    for i := range cars {
        prefixSum[i+1] = prefixSum[i] + cars[i]
        if prefixSum[i+1] != 0 {
            counter++
        }
    }
    return counter
}