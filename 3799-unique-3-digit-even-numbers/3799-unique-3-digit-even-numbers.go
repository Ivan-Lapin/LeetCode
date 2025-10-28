func totalNumbers(digits []int) int {

    freqNum := make([]int, 10)
    count := 0

    for _, digit := range digits {
        freqNum[digit] ++
    }

    for x := 100; x <= 999; x++{
        
        a := x/100
        b := (x/10)%10
        c := (x%100)%10

        if freqNum[a] != 0 {
            freqNum[a]--
            
            if freqNum[b] != 0 {
                freqNum[b]--

                if c%2==0 && freqNum[c] != 0 {
                    count++
                }

                freqNum[b]++
            }

            freqNum[a]++
        }
    }

    return count
}