func rotateString(s string, goal string) bool {
    for range len(s) {
        new_str := s[1:] + string(s[0])
        if new_str == goal {
            return true
        }
        s = new_str
    }
    return false
}