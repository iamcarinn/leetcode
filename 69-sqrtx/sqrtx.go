func ok(mid, x int) bool {
    return mid * mid <= x
}

func mySqrt(x int) int {
    ans := 0
    left := 0
    right := x
    for left <= right {
        mid := left + (right - left) / 2
        if ok(mid, x) {
            left = mid + 1
            ans = mid 
        } else {
            right = mid - 1
        }
    }

    return ans
}