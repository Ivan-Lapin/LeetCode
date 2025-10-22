func numRescueBoats(people []int, limit int) int {

    sort.Ints(people)

    count := 0

    for l, r := 0, len(people)-1; l <= r; {

        if l != r {

            if people[l] + people[r] <= limit {
                l++
            }
            r--
            count++

        } else {

            count ++
            break
            
        }
    }

    return count
}