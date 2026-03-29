func isValidSudoku(board [][]byte) bool {

	rows := [9][9]bool{}
    columns := [9][9]bool{}
    blocks := [9][9]bool{}

    for r := range 9 {
        for c := range 9{

            el := board[r][c]
            if el == '.' {
                continue
            }

            num := el - '1'
            block := (r/3)*3 + (c/3)

            if rows[r][num] || columns[c][num] || blocks[block][num] {
                return false
            }

            rows[r][num] = true
            columns[c][num] = true
            blocks[block][num] = true
        }
    }

	return true
}