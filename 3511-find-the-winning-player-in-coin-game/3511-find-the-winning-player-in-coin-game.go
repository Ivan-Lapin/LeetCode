func winningPlayer(x int, y int) string {

    try := 1

    for x >= 1 && y >= 4 {
        x--
        y -= 4
        try ++
    }

    if try%2 == 0{
        return "Alice"
    }

    return "Bob"

}