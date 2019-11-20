package libs

func GetHuntingInitSize(size int) (cap, maxCap int) {
	switch size {
	case 1:
		cap = 25
		maxCap = 25
	case 2:
		cap = 50
		maxCap = 50
	case 3:
		cap = 100
		maxCap = 100
	case 4:
		cap = 250
		maxCap = 250

	case 5:
		cap = 500
		maxCap = 500
	}
	return
}

func GetHuntingDayInc(cap, size int) int {
	_, maxCap := GetHuntingInitSize(size)
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
