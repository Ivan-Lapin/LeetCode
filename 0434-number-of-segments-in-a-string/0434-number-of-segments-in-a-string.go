func countSegments(s string) int {
    count := 0
    flag := false

    for i := 0; i < len(s); i++{
        fmt.Println(s[:i], count)
        if s[i] != 32 {
            flag = true
        }
        if (s[i] == 32 || i == len(s) - 1) && flag{
            count++
            flag = false
        }
    }
    return count
}