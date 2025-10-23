func ok(mid int, nums []int) bool {
    return nums[mid] > mid
}

func missingNumber(nums []int) int {
    sort.Ints(nums)

    ans := len(nums)
    left := 0
    right := len(nums) - 1
    for left <= right {
        mid := left + (right - left) / 2
        if ok(mid, nums) {
            right = mid - 1
            ans = mid
        } else {
            left = mid + 1
        }
    }
    return ans
}