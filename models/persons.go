package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type PersonMastery struct {
	Mastery libs.Mastery `json:"mastery"`
	Skill   float64      `json:"skill"`
}

type PersonDayAction struct {
	Action   string    `json:"action"`
	AreaType string    `json:"-"`
	AreaSize int       `json:"-"`
	AreaID   uuid.UUID `json:"-"`
	Today    int       `json:"-"`
}

type Person struct {
	ID         uuid.UUID                     `json:"id"`
	Name       string                        `json:"name"`
	Age        int                           `json:"age"`
	IsMale     bool                          `json:"is_male"`
	Chunk      uuid.UUID                     `json:"chunk_id"`
	InCh       chan string                   `json:"-"`
	Inventory  map[uuid.UUID]PersonInventory `json:"inventory"`
	Mastership []PersonMastery               `json:"mastership"`
	DayAction  PersonDayAction               `json:"day_action"`
}

var Persons []Person

func (p *Person) SetDayInc() {
	switch p.DayAction.Action {
	case "fishing":
		fishHaul := libs.CalcPersonFishingDayHaul(p.DayAction.AreaType, p.DayAction.AreaSize, p.getPersonMasterySkill(p.DayAction.Action))
		p.createHaul(fishHaul)
	case "hunting":
		huntHaul := libs.CalcPersonHuntingDayHaul(p.DayAction.AreaType, p.DayAction.AreaSize, p.getPersonMasterySkill(p.DayAction.Action))
		p.createHaul(huntHaul)
	case "food_gathering":
		fgHaul := libs.CalcPersonFGDayHaul(p.DayAction.AreaType, p.DayAction.AreaSize, p.getPersonMasterySkill(p.DayAction.Action))
		p.createHaul(fgHaul)
	case "waiting":
	}
	p.removeRottingItems()
	p.setDayMastery()

	//lakeUUID := GetRandLakeUUID()
	//SLakeMessage(lakeUUID, fmt.Sprintf("fishing|%s|%s", strconv.FormatFloat(p.Skill, 'f', -1, 64), p.ID.String()))
}

func PersonsNextDate() {
	for i := range Persons {
		Persons[i].InCh <- "next"
	}
}

func PersonsStart() {
	for i := range Persons {
		go Persons[i].PersonListener()
	}
}

func (p *Person) PersonListener() {
	for {
		com := <-p.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			p.SetDayInc()
		}
	}
}

func (p *Person) getInventory() map[uuid.UUID]PersonInventory {
	return p.Inventory
}

func (p *Person) createHaul(hauls []libs.Haul) {
	for i := range hauls {
		f := libs.GetMasteryItemByID(hauls[i].ID)
		item := getItemPool().(*Item)
		item.UUID = uuid.Must(uuid.NewV1())
		item.Name = f.Name
		item.Weight = hauls[i].Weight
		item.Quality = hauls[i].Qaulity
		item.Limit = f.LimitDay
		item.CreationDate = p.DayAction.Today
		item.ExpDate = item.CreationDate + item.Limit
		item.IsCountable = f.IsCountable
		item.Object.IsCountable = item.IsCountable
		item.Object.Name = f.Category
		Items = append(Items, item)
		p.Inventory[item.UUID] = PersonInventory{
			ID:           item.UUID,
			Name:         item.Name,
			Category:     f.Category,
			Weight:       item.Weight,
			Limit:        item.Limit,
			Quality:      item.Quality,
			CreationDate: item.CreationDate,
			ExpDate:      item.ExpDate,
			IsCountable:  f.IsCountable}
	}
}

func (p *Person) removeRottingItems() {
	for key := range p.Inventory {
		if p.Inventory[key].ExpDate < GetDate() {
			item := getItemByUUID(p.Inventory[key].ID)
			delete(p.Inventory, key)
			putItemToPool(item)
		}
	}
}

/*
func (p *Person) setFishingResult(res string) {
	var (
		hauls []FishHaul
		dM    float64
	)
	err := json.Unmarshal([]byte(res), &hauls)
	if err != nil {
		fmt.Printf("При маршалинге улова в JSON у персонажа %d произошла ошибка %s", p.ID, err)
		return
	}
	for i := range hauls {
		f := getFishByID(hauls[i].ID)
		// Считаем среднее квадратическое между редкостью рыбы и её качеством
		dM1 := math.Pow(float64(hauls[i].Qaulity*f.Rarity), 0.5)
		// Берём целое от деления массы рыбы на 1000 и добавляем 1
		dM2 := float64((hauls[i].Weight/1000)+1) / 2
		// Перемножаем 1 и 2 и делим на 10
		dM3 := (dM1 * float64(dM2)) / 10.0
		// Берём понижающий коэффициент как (100 - уровень навыка) / 100
		dM4 := (100 - p.Skill) / 100.0
		// Берём произведение 3 и 4 - получаем базовый прирост уровня навыка за конкретную пойманую рыбу
		dM5 := dM3 * dM4
		//  Рандомно добавляем к этому значению от -20% до +20% - получаем итоговый прирост уровня навыка за рыбину
		dM6 := float64(libs.GetRandInt(0, 40)-20) / 100.0
		dM7 := (dM5 + dM5*float64(dM6)) / 10
		dM += dM7
		//fmt.Printf("Персонаж %s поймал рыбу %s с весом %d и качеством %d\n", p.Name, f.Name, hauls[i].Weight, hauls[i].Qaulity)
		//fmt.Printf("Расчёт мастерства: dM1 = %f, dM2 = %f, dM3 = %f, dM4 = %f, dM5 = %f, dM6 = %f. Итоговый прирост - %f, суммарно - %f\n\n", dM1, dM2, dM3, dM4, dM5, dM6, dM7, dM)
	}

	p.Skill += dM
	p.Skill = math.Round(p.Skill*100) / 100
	if p.Skill > 100.0 {
		p.Skill = 100.0
	}
	NewEvent(fmt.Sprintf("Персонаж %s выловил %d рыбы и получил прирост навыка рыбалки на %f. Текущее значение навыка - %f", p.Name, len(hauls), dM, p.Skill))
}
*/
func PersonMessage(id uuid.UUID, text string) {
	p, err := getPersonByUUID(id)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	p.InCh <- text
}

func getPersonByUUID(id uuid.UUID) (Person, error) {
	for i := range Persons {
		if uuid.Equal(Persons[i].ID, id) {
			return Persons[i], nil
		}
	}
	return Persons[0], errors.New("Такой персонаж не найдено\n")
}

func (p *Person) getPersonMasterySkill(mastery string) float64 {
	s := 0.0
	for i := range p.Mastership {
		if p.Mastership[i].Mastery.NameID == mastery {
			s = p.Mastership[i].Skill
			break
		}
	}
	return s
}

func GetPersonInventory(param string) (inv []PersonInventory) {
	inv = make([]PersonInventory, 0, 0)
	id, err := uuid.FromString(param)
	if err != nil {
		fmt.Printf("При получении ID персонажа %s произошла ошибка %s", param, err)
		return
	}
	p, err := getPersonByUUID(id)
	if err != nil {
		fmt.Printf("При получении персонажа %s произошла ошибка %s", param, err)
		return
	}
	pi := p.getInventory()
	for k, _ := range pi {
		item := getItemByUUID(k)
		pitem := PersonInventory{
			ID:           item.UUID,
			Name:         item.Name,
			Category:     item.Object.Name,
			Weight:       item.Weight,
			Limit:        item.Limit,
			Quality:      item.Quality,
			CreationDate: item.CreationDate,
			ExpDate:      item.ExpDate,
			IsCountable:  item.Object.IsCountable}
		inv = append(inv, pitem)
	}
	return
}

func CreatePerson() {
	Persons = make([]Person, 3)
	Persons = []Person{
		Person{
			ID:        uuid.Must(uuid.NewV1()),
			Name:      "Эльсил Осландор",
			Age:       31,
			IsMale:    true,
			Chunk:     uuid.Must(uuid.FromString("36104469-81e8-4896-9790-4828c65e913c")),
			InCh:      make(chan string, 0),
			Inventory: make(map[uuid.UUID]PersonInventory),
			Mastership: []PersonMastery{
				PersonMastery{
					Mastery: libs.GetMasteryByName("fishing"),
					Skill:   25},
				PersonMastery{
					Mastery: libs.GetMasteryByName("hunting"),
					Skill:   30},
				PersonMastery{
					Mastery: libs.GetMasteryByName("food_gathering"),
					Skill:   10},
			},
			DayAction: PersonDayAction{
				Action: "waiting"}},
		Person{
			ID:        uuid.Must(uuid.NewV1()),
			Name:      "Георгил Герон",
			Age:       16,
			IsMale:    true,
			Chunk:     uuid.Must(uuid.FromString("bfff4f61-aae3-4fb6-b2f9-84e4638e270a")),
			InCh:      make(chan string, 0),
			Inventory: make(map[uuid.UUID]PersonInventory),
			Mastership: []PersonMastery{
				PersonMastery{
					Mastery: libs.GetMasteryByName("fishing"),
					Skill:   15},
				PersonMastery{
					Mastery: libs.GetMasteryByName("hunting"),
					Skill:   5},
				PersonMastery{
					Mastery: libs.GetMasteryByName("food_gathering"),
					Skill:   27},
			},
			DayAction: PersonDayAction{
				Action: "waiting"}},
		Person{
			ID:        uuid.Must(uuid.NewV1()),
			Name:      "Фарибар Тартелидил",
			Age:       54,
			IsMale:    true,
			Chunk:     uuid.Must(uuid.FromString("36104469-81e8-4896-9790-4828c65e913c")),
			InCh:      make(chan string, 0),
			Inventory: make(map[uuid.UUID]PersonInventory),
			Mastership: []PersonMastery{
				PersonMastery{
					Mastery: libs.GetMasteryByName("fishing"),
					Skill:   41},
				PersonMastery{
					Mastery: libs.GetMasteryByName("hunting"),
					Skill:   16},
				PersonMastery{
					Mastery: libs.GetMasteryByName("food_gathering"),
					Skill:   19},
			},
			DayAction: PersonDayAction{
				Action: "waiting"}}}
	for i := range Persons {
		Persons[i].setDayMastery()
	}
}

func (p *Person) setDayMastery() {
	var pm map[string]AreaMastery = make(map[string]AreaMastery)
	areaMast := GetChunckAreasMastery(p.Chunk)
	var pda PersonDayAction
	dayMasteryIndex := 0.0
	for _, m := range p.Mastership {
		for k := range areaMast {
			if k == m.Mastery.NameID {
				s := 0
				var am AreaMastery
				for i := range areaMast[k] {
					if s < areaMast[k][i].Size {
						am = areaMast[k][i]
						s = areaMast[k][i].Size
					}
				}
				pm[k] = am
			}
		}
		if dayMasteryIndex < float64(pm[m.Mastery.NameID].Size)*m.Skill {
			dayMasteryIndex = float64(pm[m.Mastery.NameID].Size) * m.Skill
			pda.Action = m.Mastery.NameID
			pda.AreaID = pm[m.Mastery.NameID].AreaID
			pda.AreaSize = pm[m.Mastery.NameID].Size
			pda.AreaType = pm[m.Mastery.NameID].Name
			if p.DayAction.Today != 0 && p.DayAction.Today == GetDate() {
				pda.Today++
			} else {
				pda.Today = GetDate()
			}

		}
	}
	p.DayAction = pda
}

func GetPersons() []Person {
	return Persons
}

/*
func getRandMale() bool {
	return libs.GetRandInt(0, 1) != 0
}

func getRandName(isMale bool) string {
	femaleName := []string{
		"Лаиммика",
		"Аимит",
		"Севикисса",
		"Севина",
		"Китеит",
		"Суэтра",
		"Беэтсена",
		"Баизжина",
		"Дефака",
		"Йизла",
		"Йивда",
		"Вафида",
		"Ларила",
		"Шевирисса",
		"Саадгена",
		"Суфиза",
		"Сиоркисса",
		"Меидин",
		"Доимас",
		"Диитва",
		"Сафаника",
		"Уваника",
		"Йдирисса",
		"Сеорда",
		"Луадет"}
	femaleSurname := []string{
		"Суивра",
		"Ваитсена",
		"Деривена",
		"Сафада",
		"Злаоркисса",
		"Иилам",
		"Секиам",
		"Дореит",
		"Миивас",
		"Вавигена",
		"Викиза",
		"Сеилам",
		"Довива",
		"Саадника",
		"Судирисса",
		"Лаиммика",
		"Златеит",
		"Деэтвена",
		"Дивива",
		"Мадиам",
		"Бешина",
		"Мадиим",
		"Сирижина",
		"Бефает",
		"Шиорда"}
	maleName := []string{
		"Беронлас",
		"Араланвир",
		"Белек",
		"Экоркар",
		"Дилекфор",
		"Осрин",
		"Фарилеб",
		"Эльсил",
		"Фарадор",
		"Георгил",
		"Дирорн",
		"Тулан",
		"Тартелидил",
		"Динакил",
		"Элховал",
		"Мерендил",
		"Экор",
		"Арарадур",
		"Ботелитар",
		"Экланлион",
		"Тарон",
		"Валангил",
		"Эльлетур",
		"Фарибар",
		"Харин"}
	maleSurname := []string{
		"Кинадур",
		"Туормир",
		"Тарнанон",
		"Герон",
		"Элрион",
		"Борен",
		"Денелеб",
		"Баоддур",
		"Эльтеливир",
		"Кивирил",
		"Оснадор",
		"Меагбар",
		"Элрилор",
		"Халелеб",
		"Элдор",
		"Вахогил",
		"Банадур",
		"Тухогил",
		"Басилдил",
		"Меланкар",
		"Тарондур",
		"Араланвир",
		"Эллекдил",
		"Эльлан",
		"Месил"}
	if isMale {
		return maleName[libs.GetRandInt(0, len(maleName)-1)] + " " + maleSurname[libs.GetRandInt(0, len(maleSurname)-1)]
	}
	return femaleName[libs.GetRandInt(0, len(femaleName)-1)] + " " + femaleSurname[libs.GetRandInt(0, len(femaleSurname)-1)]
}
*/
