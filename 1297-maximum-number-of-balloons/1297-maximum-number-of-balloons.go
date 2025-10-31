func maxNumberOfBalloons(text string) int {

    if len(text) < 7 {
		return 0
	}

    freqChar := make(map[rune]int)
    count := 0

    for i := 0; i < len(text); i++{
        freqChar[rune(text[i])] ++
    }

    var def func(mapa map[rune]int) bool 

    def = func(mapa map[rune]int) bool {
        for _, ch := range []rune{'b', 'a', 'l', 'l', 'o', 'o', 'n'} {
            if val, exist := mapa[ch]; exist && val == 0 || !exist{
                return false
            } else {
                mapa[ch] --
            }
        }
        return true
    }

    for def(freqChar) {
        count++
    }

    return count
}