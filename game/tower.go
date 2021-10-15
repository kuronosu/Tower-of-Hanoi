package game

type Tower []int

func (tower Tower) IsEmpty() bool {
	return len(tower) == 0
}

func (tower *Tower) Push(str int) error {
	*tower = append(*tower, str)
	return nil
}

func (tower *Tower) Pop() (int, bool) {
	if tower.IsEmpty() {
		return 0, false
	} else {
		index := len(*tower) - 1
		element := (*tower)[index]
		*tower = (*tower)[:index]
		return element, true
	}
}

func (tower Tower) Eq(other Tower) bool {
	if len(tower) != len(other) {
		return false
	}
	for i, v := range tower {
		if v != other[i] {
			return false
		}
	}
	return true
}

func (tower Tower) IsValid() bool {
	if tower.IsEmpty() {
		return true
	}
	prev := tower[0]
	for _, current := range tower[1:] {
		if current > prev {
			return false
		}
		prev = current
	}
	return true
}

func (tower Tower) Clone() *Tower {
	tmp := &Tower{}
	for _, v := range tower {
		tmp.Push(v)
	}
	return tmp
}

func (tower Tower) GetDisc() int {
	if len(tower) <= 0 {
		return -1
	}
	return tower[len(tower)-1]
}
