package models

import "strconv"

type Lake struct {
	ID          int
	Size        int
	MaxCapacity int
	Capacity    int
	DayInc      int
	InCh        chan string `json:"-"`
}

func (l *Lake) SetDayInc() {
	l.DayInc = int(l.MaxCapacity / 100)
	l.Capacity += l.DayInc
	if l.Capacity > l.MaxCapacity {
		l.DayInc -= (l.Capacity - l.MaxCapacity)
		l.Capacity = l.MaxCapacity
	}
	formatString := "Озеро: " + strconv.Itoa(l.ID) + ". Прирост за день - " + strconv.Itoa(l.DayInc) + ", текущее значение - " + strconv.Itoa(l.Capacity)
	NewEvent(formatString)
}

func (l *Lake) LakeNextDate() {
	for {
		com := <-l.InCh
		if com == "next" {
			l.SetDayInc()
		}
	}
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
			0,
			make(chan string, 0)}
	}
}

func LakesStart() {
	for l := range Lakes {
		go Lakes[l].LakeNextDate()
	}
}

func LakesNextDate() {
	for l := range Lakes {
		Lakes[l].InCh <- "next"
	}
}

func GetLakes() []Lake {
	return Lakes
}
