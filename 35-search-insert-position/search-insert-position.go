func searchInsert(nums []int, target int) int {
    ans := 0
    if nums[0] > target {
        return 0
    }

    for i := range nums {
        if nums[i] < target {
            ans = i
        } else if nums[i] == target {
            ans = i - 1
        }
    }
    return ans + 1
}
