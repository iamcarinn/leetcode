func shiftingLetters(s string, shifts [][]int) string {
    n := len(s)
    directions := make([]int, n+1)

    for i := range shifts {
        start, end, direction := shifts[i][0], shifts[i][1], shifts[i][2]
        if direction == 1 {
            directions[start]++
            directions[end+1]--
        } else {
            directions[start]--
            directions[end+1]++
        }
    }

    result := []byte(s)
    curr_shift := 0
    for i := 0; i < n; i++ {
        curr_shift += directions[i]
        shifted := (int(s[i]-'a') + curr_shift) % 26

        if shifted < 0 {
            shifted += 26
        }
        result[i] = byte(shifted) + 'a'
    }

    return string(result)
}