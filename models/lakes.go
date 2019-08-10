package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Lake struct {
	ID          int
	Size        int
	MaxCapacity int
	Capacity    int
	DayInc      int
	UUID        uuid.UUID   `json:"-"`
	InCh        chan string `json:"-"`
}

func (l *Lake) SetDayInc() {
	l.DayInc = int(l.MaxCapacity / 100)
	l.Capacity += l.DayInc
	if l.Capacity > l.MaxCapacity {
		l.DayInc -= (l.Capacity - l.MaxCapacity)
		l.Capacity = l.MaxCapacity
	}
	NewEvent(
		fmt.Sprintf("Озеро: %s. Прирост за день - %s, текущее значение - %s",
			strconv.Itoa(l.ID),
			strconv.Itoa(l.DayInc),
			strconv.Itoa(l.Capacity)))
}

func (l *Lake) LakeListener() {
	for {
		com := <-l.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go l.SetDayInc()
		case "fishing":
			go l.calcFishingResult(params[1], params[2])
		}
	}
}

var Lakes []Lake

func (l *Lake) calcFishingResult(skill, personId string) {
	s, err := strconv.ParseFloat(skill, 64)
	if err != nil {
		fmt.Printf("При получении навыка рыбалки произошла ошибка %s у персонажа %s", err, personId)
		s = 0
	}
	res := (int)(s*2) * l.Size
	if res > l.Capacity {
		res = l.Capacity
	}
	l.Capacity -= res
	personUUID, err := uuid.FromString(personId)
	if err != nil {
		fmt.Printf("При получении UUID рыбака произошла ошибка %s у персонажа %s", err, personId)
	} else {
		PersonMessage(personUUID, fmt.Sprintf("fishing|%s|%s", strconv.Itoa(res), l.UUID.String()))
	}
	NewEvent(fmt.Sprintf("Из озера %s выловили рыбин: %s, текущее значение - %s",
		strconv.Itoa(l.ID), strconv.Itoa(res), strconv.Itoa(l.Capacity)))
}

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
			uuid.Must(uuid.NewV1()),
			make(chan string, 0)}
	}
}

func LakesStart() {
	for l := range Lakes {
		go Lakes[l].LakeListener()
	}
}

func LakesNextDate() {
	for l := range Lakes {
		Lakes[l].InCh <- "next"
	}
}

func LakeMessage(id uuid.UUID, text string) {
	l, err := getLakeByUUID(id)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	l.InCh <- text
}

func getLakeByUUID(id uuid.UUID) (Lake, error) {
	for i := range Lakes {
		if uuid.Equal(Lakes[i].UUID, id) {
			return Lakes[i], nil
		}
	}
	return Lakes[0], errors.New("Такое озеро не найдено\n")
}

func GetLakes() []Lake {
	return Lakes
}

func GetRandLakeUUID() uuid.UUID {
	return Lakes[GetRandInt(0, len(Lakes)-1)].UUID
}
