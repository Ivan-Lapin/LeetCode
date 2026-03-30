type RandomizedSet struct {
	repo  map[int]int
	slice []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		repo:  make(map[int]int),
		slice: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.repo[val]; exist {
		return false
	}
	this.slice = append(this.slice, val)
    this.repo[val] = len(this.slice)-1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, exist := this.repo[val]; exist {
        last_el := this.slice[len(this.slice)-1]
        this.repo[last_el] = index
        this.slice[index] = last_el
        this.slice = this.slice[:len(this.slice)-1]
        delete(this.repo, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
    return this.slice[rand.Intn(len(this.slice))]
}