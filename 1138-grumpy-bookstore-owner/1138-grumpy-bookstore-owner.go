func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    realy := 0
    pozitive := 0
    not_grumpy := 0

    for i := 0; i < minutes; i++{
        if grumpy[i] == 0 {
            realy += customers[i]
            not_grumpy += customers[i]
        }
        pozitive += customers[i]
    }

    if minutes == len(customers){
        return pozitive
    }

    maxPozitive := pozitive - realy

    for i := minutes; i < len(customers); i++{
        
        if grumpy[i-minutes] == 0 {
            realy -= customers[i-minutes]
        }
        pozitive -= customers[i-minutes]

        if grumpy[i] == 0 {
            realy += customers[i]
            not_grumpy += customers[i]
        }
        pozitive += customers[i]

        if maxPozitive < pozitive - realy {
            maxPozitive = pozitive - realy
        }
    }

    return not_grumpy + maxPozitive
}