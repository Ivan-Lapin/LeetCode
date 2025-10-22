func alternateDigitSum(n int) int {

    result := 0

    flag := false

    str := strconv.Itoa(n)

    if len(str) % 2 != 0{
        flag = true
    }

    if flag {
        for i := 0; n > 0; {
        
            num := n%10
            if i % 2 == 0{
                result += num
            } else {
                result -= num
            }
            n /= 10
            i++
            fmt.Println(num, result)
        }
    } else {
        for i := 0; n > 0; {
        
            num := n%10
            if i % 2 != 0{
                result += num
            } else {
                result -= num
            }
            n /= 10
            i++
            fmt.Println(num, result)
        }
    }

    

    return result
}