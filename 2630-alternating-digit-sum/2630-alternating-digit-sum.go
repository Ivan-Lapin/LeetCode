func alternateDigitSum(n int) int {

    result := 0

    number := strconv.Itoa(n)

    for i := 0; i < len(number); i++{
        if i%2 == 0 {
            result += int(number[i]-'0')
        } else {
            result -= int(number[i]-'0')
        }
    }

    

    return result
}