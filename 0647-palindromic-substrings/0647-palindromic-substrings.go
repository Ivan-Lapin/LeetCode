func countSubstrings(s string) int {
    
    count := 0
    for i := 0; i < len(s); i++{
        count++
        for j := 1; i-j >=0 && i+j < len(s); j++ {
            if s[i-j] == s[i+j] {
                count++
            } else {
                break
            }
        }
    }

    for l, r := 0, 1; r < len(s); l, r = l+1, r+1{
        if s[l] == s[r] {
            count++
            for j := 1; l-j >=0 && r+j < len(s); j++ {
                if s[l-j] == s[r+j] {
                    count++
                } else {
                    break
                }
            }
        }
    }    

    return count
}