func restoreString(s string, indices []int) string {
    arr := make([]byte, len(indices))
    for i, val := range indices{
        arr[val] = s[i]
    }
    return string(arr)
}