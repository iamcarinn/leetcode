func longestCommonPrefix(strs []string) string {
    
    // ищем максимальную длину префикса по минимульной длине слова
    maxLenP := len(strs[0])
    for _, s := range strs[1:] {
        if len([]rune(s)) < maxLenP{
            maxLenP = len(s)
        }
    }

    bestPref := ""
    // идем увеличивая размер префикса от 1 до maxLenP
    for lenPref := 1; lenPref <= maxLenP; lenPref++ {
        // идем циклом по другим словам
        prefix := strs[0][:lenPref]
        ok := true
        for i := 1; i < len(strs); i++ {
            if !strings.HasPrefix(strs[i], prefix) {
                ok = false
                break
            }
            
        }
        if ok {
            bestPref = prefix
        }
    }
    return bestPref
}
