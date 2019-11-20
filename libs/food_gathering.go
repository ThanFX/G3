package libs

func GetFoodGatheringInitSize(size int) (cap, maxCap int) {
	switch size {
	case 1:
		cap = 25000
		maxCap = 25000
	case 2:
		cap = 50000
		maxCap = 50000
	case 3:
		cap = 100000
		maxCap = 100000
	case 4:
		cap = 250000
		maxCap = 250000
	case 5:
		cap = 500000
		maxCap = 500000
	}
	return
}

func GetFoodGatheringDayInc(cap, size int) int {
	_, maxCap := GetFoodGatheringInitSize(size)
	dayInc := int((maxCap - cap) / 100)
	if dayInc < 1 {
		dayInc = 1
	}
	cap += dayInc
	if cap > maxCap {
		dayInc -= (cap - maxCap)
		cap = maxCap
	}
	return cap
}
