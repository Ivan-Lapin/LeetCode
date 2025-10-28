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
            result = fmt.Sprint(num%7) + result
            num = div
        }
    }

    if num != 0 {
        result = fmt.Sprint(num) + result
    }

    return first_part + result

}