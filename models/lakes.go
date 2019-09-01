package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type FishHaul struct {
	Name    string
	Weight  int
	Qaulity string
}

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

func (l *Lake) setDayInc() {
	l.DayInc = int((l.MaxCapacity - l.Capacity) / 100)
	if l.DayInc < 1 {
		l.DayInc = 1
	}
	l.Capacity += l.DayInc
	if l.Capacity > l.MaxCapacity {
		l.DayInc -= (l.Capacity - l.MaxCapacity)
		l.Capacity = l.MaxCapacity
	}
	NewEvent(
		fmt.Sprintf("В озере %s за день родилось %s рыбы. Всего сейчас %s рыбы.", strconv.Itoa(l.ID), strconv.Itoa(l.DayInc), strconv.Itoa(l.Capacity)))
}

func (l *Lake) lakeListener() {
	for {
		com := <-l.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go l.setDayInc()
		case "fishing":
			go l.calcFishingResult(params[1], params[2])
		}
	}
}

var (
	Lakes   []Lake
	Fishes  []Fish
	DB      *sql.DB
	Quality = [5]string{"Обычная", "Хорошая", "Отличная", "Превосходная", "Идеальная"}
)

func getMasteryFunc(x float64) float64 {
	res := math.Round((math.Pow(1.055, x)/2)*100) / 100
	if res < 1 {
		res = 1
	} else if res > 100 {
		res = 100
	}
	return res
}

func getFishNameForAreaSize(areaSize, rarity int) string {
	var fishes []Fish
	//fmt.Printf("Ищем рыбу в озере размером %d и с редкостью %d\n", areaSize, rarity)
	for i := range Fishes {
		if Fishes[i].Area <= areaSize && Fishes[i].Rarity <= rarity {
			//fmt.Println(Fishes[i])
			fishes = append(fishes, Fishes[i])
		}
	}
	//fmt.Println(fishes)
	return fishes[GetRandInt(0, len(fishes)-1)].Name
}

func (l *Lake) calcFishingResult(skill, personId string) {
	s, err := strconv.ParseFloat(skill, 64)
	if err != nil {
		fmt.Printf("При получении навыка рыбалки произошла ошибка %s у персонажа %s", err, personId)
		s = 0
	}
	var hauls []FishHaul
	mastery := getMasteryFunc(s)
	// 32 тика: 8 часов * 4 скорость удочки
	for i := 0; i < 32; i++ {
		// шанс на улов
		f := float64(0.25) + (mastery / 100)
		if rand.Float64() < f && l.Capacity > 2 {
			// Поймали рыбу, теперь нужно определить её параметры
			// Считаем максимальный улов (чем выше рандом, тем выше вес)
			randWeight := GetRandInt(1, int(math.Floor(mastery)))
			var haul FishHaul
			wf := float64(l.MaxCapacity/50) * float64(randWeight/10)
			if wf < 100 {
				wf = 100
			} else {
				wf = math.Floor(wf/100) * 100
			}
			//fmt.Printf("На тике №%d мастерство %f, поймали рыбину весом %f грамм. Шанс на вес - %d\n", i+1, mastery, wf, randWeight)
			haul.Weight = int(wf)

			// Теперь определяем вид пойманой рыбы
			var fishRarity = 1
			for dice := 0; dice < int(mastery/10)+1; dice++ {
				rar := 1
				chance := GetRandInt(0, 1000000)
				switch {
				case chance > 990000:
					rar = 5
				case chance > 960000:
					rar = 4
				case chance > 900000:
					rar = 3
				case chance > 600000:
					rar = 2
				}
				//fmt.Printf("Выпал шанс %d, рыба уровня %d\n", chance, rar)
				if rar > fishRarity {
					fishRarity = rar
				}
			}
			haul.Name = getFishNameForAreaSize(l.Size, fishRarity)

			// Определяем качество пойманной рыбы
			var fishQuality = 1
			for dice := 0; dice < int(mastery/10)+1; dice++ {
				qual := 1
				chance := GetRandInt(0, 10000000)
				switch {
				case chance > 9980000:
					qual = 5
				case chance > 9900000:
					qual = 4
				case chance > 9500000:
					qual = 3
				case chance > 8000000:
					qual = 2
				}
				//fmt.Printf("Выпал шанс %d, рыба уровня %d\n", chance, rar)
				if qual > fishQuality {
					fishQuality = qual
				}
			}
			haul.Qaulity = Quality[fishQuality-1]
			hauls = append(hauls, haul)
			l.Capacity--
			//fmt.Printf("На тике №%d поймали рыбу \"%s\" весом %d грамм с качеством \"%s\"\n", i+1, haul.Name, haul.Weight, haul.Qaulity)
		} else {
			//fmt.Printf("На тике №%d нихуя не поймали\n", i+1)
		}
	}

	personUUID, err := uuid.FromString(personId)
	if err != nil {
		fmt.Printf("При получении UUID рыбака произошла ошибка %s у персонажа %s", err, personId)
	} else {
		PersonMessage(personUUID, fmt.Sprintf("fishing|%s|%s", strconv.Itoa(len(hauls)), l.UUID.String()))
	}

	for i := range hauls {
		NewEvent(fmt.Sprintf("Поймали рыбу \"%s\" весом %d грамм с качеством \"%s\"", hauls[i].Name, hauls[i].Weight, hauls[i].Qaulity))
	}

	id := strconv.Itoa(l.ID)
	r := strconv.Itoa(len(hauls))
	NewEvent(fmt.Sprintf("Из озера %s выловили %s рыбы. Осталось %s рыбы.", id, r, strconv.Itoa(l.Capacity)))
}

func CreateLakes(count int) {
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

func getFishByLakeSize(size int) []Fish {
	var f []Fish
	for i := range Fishes {
		if Fishes[i].Area <= size {
			f = append(f, Fishes[i])
		}
	}
	return f
}

func readFishCatalog() {
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
		go Lakes[l].lakeListener()
	}
	readFishCatalog()
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
