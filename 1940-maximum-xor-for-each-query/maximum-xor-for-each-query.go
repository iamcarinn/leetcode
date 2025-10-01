func getMaximumXor(nums []int, maximumBit int) []int {
    prefix_xor := make([]int, len(nums))
    prefix_xor[0] = nums[0]

    for i := 1; i < len(nums); i++ {
        prefix_xor[i] = prefix_xor[i-1] ^ nums[i]
    }

    maxNum := (1 << maximumBit) - 1

    res := make([]int, len(nums))
    for i := range res {
        res[i] = prefix_xor[len(nums) - i - 1] ^ maxNum
    }

    return  res
}