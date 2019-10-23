package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/ThanFX/G3/libs"

	uuid "github.com/satori/go.uuid"
)

type FishHaul struct {
	ID      int
	Weight  int
	Qaulity int
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

var (
	Lakes  []Lake
	Fishes []Fish
	//DB      *sql.DB
	Quality = [5]string{"Обычная", "Хорошая", "Отличная", "Превосходная", "Идеальная"}
)

/*
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
*/

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
	return Lakes[libs.GetRandInt(0, len(Lakes)-1)].UUID
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

func getFishNameForAreaSize(areaSize, rarity int) int {
	var fishes []Fish
	for i := range Fishes {
		if Fishes[i].Area <= areaSize && Fishes[i].Rarity <= rarity {
			fishes = append(fishes, Fishes[i])
		}
	}
	return fishes[libs.GetRandInt(0, len(fishes)-1)].ID
}

func getFishByID(id int) Fish {
	var f Fish
	for i := range Fishes {
		if Fishes[i].ID == id {
			f = Fishes[i]
			break
		}
	}
	return f
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

func (l *Lake) calcFishingResult(skill, personId string) {
	s, err := strconv.ParseFloat(skill, 64)
	if err != nil {
		fmt.Printf("При получении навыка рыбалки произошла ошибка %s у персонажа %s", err, personId)
		s = 0
	}
	var hauls []FishHaul
	mastery := getMasteryFunc(s)
	//NewEvent(fmt.Sprintf("Уровень навыка - %f, мастерство - %f", s, mastery))
	// 32 тика: 8 часов * 4 скорость удочки
	for i := 0; i < 32; i++ {
		// шанс на улов
		f := float64(0.1) + (mastery / 100)
		if rand.Float64() < f && l.Capacity > 2 {
			// Поймали рыбу, теперь нужно определить её параметры
			// Считаем максимальный улов (чем выше рандом, тем выше вес)
			randWeight := libs.GetRandInt(1, int(math.Floor(mastery)))
			var haul FishHaul
			// Максимальная масса рыбы, которую можно выловить с таким уровнем навыка и в таком озере
			wfMax := float64(l.MaxCapacity/50) * float64(randWeight/10)
			if wfMax < 100 {
				wfMax = 100
			} else {
				wfMax = math.Floor(wfMax/100) * 100
			}
			wf := libs.GetRandInt(1, int(wfMax)/100) * 100
			//fmt.Printf("На тике №%d мастерство %f, поймали рыбину весом %f грамм. Шанс на вес - %d\n", i+1, mastery, wf, randWeight)
			haul.Weight = int(wf)

			// Теперь определяем вид пойманой рыбы
			var fishRarity = 1
			for dice := 0; dice < int(mastery/10)+1; dice++ {
				rar := 1
				chance := libs.GetRandInt(0, 1000000)
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
			haul.ID = getFishNameForAreaSize(l.Size, fishRarity)

			// Определяем качество пойманной рыбы
			var fishQuality = 1
			for dice := 0; dice < int(mastery/10)+1; dice++ {
				qual := 1
				chance := libs.GetRandInt(0, 10000000)
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
			haul.Qaulity = fishQuality
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
		haulJSON, err := json.Marshal(hauls)
		if err != nil {
			fmt.Printf("При маршалинге улова в JSON у персонажа %s произошла ошибка %s", personId, err)
		}
		PersonMessage(personUUID, fmt.Sprintf("fishing|%s", string(haulJSON)))
	}

	id := strconv.Itoa(l.ID)
	r := strconv.Itoa(len(hauls))
	NewEvent(fmt.Sprintf("Из озера %s выловили %s рыбы. Осталось %s рыбы.", id, r, strconv.Itoa(l.Capacity)))
}

func CreateLakes(count int) {
	Lakes = make([]Lake, count)
	for i := range Lakes {
		size := libs.GetRandInt(1, 5)
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

func LakesStart() {
	for l := range Lakes {
		go Lakes[l].lakeListener()
	}
	//readFishCatalog()
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
