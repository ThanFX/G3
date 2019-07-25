package models

type Lake struct {
	ID          int
	Size        int
	MaxCapacity int
	Capacity    int
	DayInc      int
}

var Lakes []Lake

func CreateLakes(count int) {
	Lakes = make([]Lake, count)
	for i := range Lakes {
		size := GetRandInt(1, 3)
		var maxCap, cap int
		switch size {
		case 1:
			maxCap = 5000
			cap = 2500
		case 2:
			maxCap = 10000
			cap = 5000
		case 3:
			maxCap = 20000
			cap = 10000
		}
		Lakes[i] = Lake{
			i + 1,
			size,
			maxCap,
			cap,
			0}
	}
}

func GetLakes() []Lake {
	return Lakes
}
