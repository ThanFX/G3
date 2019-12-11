package libs

import (
	"database/sql"
	"log"
	"math"
	"strings"
)

type Haul struct {
	ID      int
	Weight  int
	Qaulity int
}

type Mastery struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	NameID   string `json:"id"`
	MinValue int    `json:"-"`
	MaxValue int    `json:"-"`
}

type MasteryItem struct {
	ID          int
	Name        string
	Mastery     string
	Category    string
	Ingredient  string
	Areas       []string
	AreaSize    int
	Rarity      int
	Min         int
	Max         int
	IsCountable bool
	IsLiquid    bool
	LimitDay    int
}

var (
	DB           *sql.DB
	MasteryItems []MasteryItem
	Quality                         = [5]string{"Обычное", "Хорошее", "Отличное", "Превосходное", "Идеальное"}
	Masterships  map[string]Mastery = make(map[string]Mastery)
)

func GetMasterships() map[string]Mastery {
	return Masterships
}

func GetMasteryByName(name string) Mastery {
	return Masterships[name]
}

func ReadMastershipsCatalog() {
	var m Mastery
	rows, err := DB.Query("select * from masterships")
	if err != nil {
		log.Fatalf("Ошибка получения профессий из БД: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&m.ID, &m.Name, &m.NameID, &m.MinValue, &m.MaxValue)
		if err != nil {
			log.Fatal("ошибка парсинга записи мастерства: ", err)
		}
		Masterships[m.NameID] = m
	}
}

func ReadMateryItemsCatalog() {
	rows, err := DB.Query("select * from mastery_items")
	if err != nil {
		log.Fatalf("Ошибка получения предметов из БД: %s", err)
	}
	defer rows.Close()

	var mi MasteryItem
	var a string
	for rows.Next() {
		err = rows.Scan(
			&mi.ID,
			&mi.Mastery,
			&mi.Category,
			&mi.Ingredient,
			&mi.Name,
			&mi.Rarity,
			&a,
			&mi.AreaSize,
			&mi.Min,
			&mi.Max,
			&mi.IsCountable,
			&mi.IsLiquid,
			&mi.LimitDay)
		if err != nil {
			log.Fatal("ошибка парсинга записи о предмете: ", err)
		}
		mi.Areas = strings.Split(a, ",")
		MasteryItems = append(MasteryItems, mi)
	}
}

func (mi *MasteryItem) IsAreaContains(areaType string) bool {
	for _, s := range mi.Areas {
		if areaType == s {
			return true
		}
	}
	return false
}

func getRandMasteryItemIDFromArea(mastery, areaType string, areaSize, rarity int) MasteryItem {
	var fishes []MasteryItem
	for i := range MasteryItems {
		if mastery == MasteryItems[i].Mastery &&
			MasteryItems[i].IsAreaContains(areaType) &&
			MasteryItems[i].AreaSize <= areaSize &&
			MasteryItems[i].Rarity <= rarity {
			fishes = append(fishes, MasteryItems[i])
		}
	}
	return fishes[GetRandInt(0, len(fishes)-1)]
}

func GetMasteryItemByID(id int) MasteryItem {
	var f MasteryItem
	for i := range MasteryItems {
		if MasteryItems[i].ID == id {
			f = MasteryItems[i]
			break
		}
	}
	return f
}

func getMasteryFunc(x float64) float64 {
	res := -0.0000001329*math.Pow(x, 5) + 0.0000279284*math.Pow(x, 4) - 0.0017605406*math.Pow(x, 3) +
		0.0339750737*math.Pow(x, 2) + 0.599870964*x + 1.1256425921
	if res < 1 {
		res = 1
	} else if res > 100 {
		res = 100
	}
	return res
}
