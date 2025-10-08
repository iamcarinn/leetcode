func maximumOr(nums []int, k int) int64 {
    n := len(nums)
	multipler := math.Pow(float64(2), float64(k))
	maxVal := 0

    prefix_or := make([]int, n)
    suffix_or := make([]int, n)

    for i := 0; i < n-1; i++ {
        prefix_or[i+1] = prefix_or[i] | nums[i]
    }


    for i := n-2; i >= 0; i-- {
        suffix_or[i] = suffix_or[i + 1] | nums[i + 1]
    }

    for i := 0; i < len(nums); i++ {
		maxVal = max(maxVal, prefix_or[i]|nums[i]*int(multipler)|suffix_or[i])
	}


    return int64(maxVal)
    
}
