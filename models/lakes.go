package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

type Fish struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Rarity  int    `json:"rarity"`
	Area    int    `json:"area"`
	IsLake  bool   `json:"is_lake"`
	IsRiver bool   `json:"is_river"`
}

func (l *Lake) SetDayInc() {
	l.DayInc = int(l.MaxCapacity / 100)
	l.Capacity += l.DayInc
	if l.Capacity > l.MaxCapacity {
		l.DayInc -= (l.Capacity - l.MaxCapacity)
		l.Capacity = l.MaxCapacity
	}
	NewEvent(
		fmt.Sprintf("В озере %s за день родилось %s рыбы. Всего сейчас %s рыбы.", strconv.Itoa(l.ID), strconv.Itoa(l.DayInc), strconv.Itoa(l.Capacity)))
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

var (
	Lakes  []Lake
	Fishes []Fish
	DB     *sql.DB
)

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
	id := strconv.Itoa(l.ID)
	r := strconv.Itoa(res)
	NewEvent(fmt.Sprintf("Из озера %s выловили %s рыбы. Осталось %s рыбы.", id, r, strconv.Itoa(l.Capacity)))
}

func CreateLakes(count int) {
	GetFishCatalog()
	Lakes = make([]Lake, count)
	for i := range Lakes {
		size := GetRandInt(1, 5)
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
		case 4:
			maxCap = 50000
			cap = 25000

		case 5:
			maxCap = 100000
			cap = 50000
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

func GetFishCatalog() {
	rows, err := DB.Query("select * from fishes")
	if err != nil {
		log.Fatalf("Ошибка получения рыб из БД: %s", err)
	}
	defer rows.Close()

	var f Fish
	for rows.Next() {
		err = rows.Scan(&f.ID, &f.Name, &f.Rarity, &f.IsLake, &f.IsRiver, &f.Area)
		if err != nil {
			log.Fatal("ошибка парсинга записи о рыбе: ", err)
		}
		Fishes = append(Fishes, f)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
