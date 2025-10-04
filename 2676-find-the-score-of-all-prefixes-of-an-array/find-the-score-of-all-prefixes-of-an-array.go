func findPrefixScore(nums []int) []int64 {
    max_prefix := 0
    sum := 0
    var conver []int64

    for i := range nums {
        max_prefix = max(max_prefix, nums[i])
        sum = sum + max_prefix + nums[i]
        conver = append(conver, int64(sum))
    }
    return conver
}