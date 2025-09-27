func largestAltitude(gain []int) int {
    size := len(gain)
    prefixSum := make([]int, size+1)
    result := prefixSum[0]

    for i := range gain {
        prefixSum[i+1] = prefixSum[i] + gain[i]
        if result < prefixSum[i+1] {
            result = prefixSum[i+1]
        }
    }
    return result
}