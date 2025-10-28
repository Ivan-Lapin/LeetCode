func convertToBase7(num int) string {

    if num == 0 {
        return "0"
    }

    result := ""
    first_part := ""

    if num < 0 {
        num = -num
        first_part = "-"
    }

    div := num
    for ; div != 0; div--{
        if 7 * div <= num {
            temp := strconv.Itoa(num%7)
            result = temp + result
            num = div
        }
    }

    if num != 0 {
        result = strconv.Itoa(num) + result
    }

    return first_part + result

}