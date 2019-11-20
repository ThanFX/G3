package libs

import (
	"database/sql"
	"log"
)

type Fish struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Rarity  int    `json:"rarity"`
	Area    int    `json:"area"`
	IsLake  bool   `json:"is_lake"`
	IsRiver bool   `json:"is_river"`
}

type FishHaul struct {
	ID      int
	Weight  int
	Qaulity int
}

var (
	FishQuality = [5]string{"Обычная", "Хорошая", "Отличная", "Превосходная", "Идеальная"}
	Fishes      []Fish
	DB          *sql.DB
)

func ReadFishCatalog() {
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

func GetFishByAreaSize(size int) []Fish {
	var f []Fish
	for i := range Fishes {
		if Fishes[i].Area <= size {
			f = append(f, Fishes[i])
		}
	}
	return f
}

func GetFishNameForAreaSize(areaSize, rarity int) int {
	var fishes []Fish
	for i := range Fishes {
		if Fishes[i].Area <= areaSize && Fishes[i].Rarity <= rarity {
			fishes = append(fishes, Fishes[i])
		}
	}
	return fishes[GetRandInt(0, len(fishes)-1)].ID
}

func GetFishByID(id int) Fish {
	var f Fish
	for i := range Fishes {
		if Fishes[i].ID == id {
			f = Fishes[i]
			break
		}
	}
	return f
}

func GetFishingInitSize(size int) (cap, maxCap int) {
	switch size {
	case 1:
		cap = 5000
		maxCap = 5000
	case 2:
		cap = 10000
		maxCap = 10000
	case 3:
		cap = 20000
		maxCap = 20000
	case 4:
		cap = 50000
		maxCap = 50000
	case 5:
		cap = 100000
		maxCap = 100000
	}
	return
}

func GetFishingDayInc(cap, size int) int {
	_, maxCap := GetFishingInitSize(size)
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
