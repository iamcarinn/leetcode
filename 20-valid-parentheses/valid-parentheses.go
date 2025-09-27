func isValid(s string) bool {
    res := true
    mymap := map[rune]int {'(': 1, '[': 2, '{': 3, ')': -1, ']': -2, '}': -3}
    myqueue := make([]int, 0, len(s)/2)

    for _, ch := range s {
        if mymap[ch] > 0 {
            myqueue = append(myqueue, mymap[ch])
        } else {
            if len(myqueue) == 0 {
                res = false
                break
            }
            if mymap[ch] + myqueue[len(myqueue) - 1] != 0 {
                res = false
                break
            } else {
                myqueue = myqueue[:len(myqueue) - 1]
            }
        }
    }
    if len(myqueue) != 0 {
        res = false
    }
    return res
}
