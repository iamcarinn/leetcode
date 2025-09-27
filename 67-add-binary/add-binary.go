func addBinary(a string, b string) string {
    if len(a) > len(b) {
        a, b = b, a
    }
    str_zero := strings.Repeat("0", len(b) - len(a))
    a = str_zero + a

    res := ""
    add := 0
    for i := len(a) - 1; i >= 0; i-- {
        sum := int(a[i]-'0') + int(b[i]-'0') + add
        res = string(sum % 2 + '0') + res
        add = sum / 2
    }

    if add == 1 {
        res = "1" + res
    }

    return res
}
