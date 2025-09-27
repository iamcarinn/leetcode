func isPalindrome(x int) bool {
  res := true

  if x < 0 {
    res = false
  }

  str_x := strconv.Itoa(x)
  len_x := len(str_x)
  for i := 0; i < len_x/2; i++ {
        if str_x[i] != str_x[len_x-1-i] {
            res = false
        }
    }

  return res
}
