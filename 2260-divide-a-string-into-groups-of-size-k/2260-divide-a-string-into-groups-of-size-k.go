func divideString(s string, k int, fill byte) []string {
    result := make([]string, (len(s)/k)+1)
    index := 0
    for i := 0; i < len(s); i += k {
        if i+k > len(s) {
            result[index] = s[i:]
            for j := k - len(result[index]); j > 0; j--{
                result[index] += string(fill)
            }
        } else {
            result[index] = s[i:i+k]
        }
        index++
    }
    if len(result[len(result)-1]) == 0{
        return result[:len(result)-1]
    }
    return result

}