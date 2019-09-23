package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID        uuid.UUID
	Name      string
	Birth     int
	IsMale    bool
	Chunk     int
	Skill     float64
	InCh      chan string                 `json:"-"`
	Inventory map[uuid.UUID]InventoryItem `json:"-"`
}

var Persons []Person

func (p *Person) SetDayInc() {
	lakeUUID := GetRandLakeUUID()
	LakeMessage(lakeUUID, fmt.Sprintf("fishing|%s|%s", strconv.FormatFloat(p.Skill, 'f', -1, 64), p.ID.String()))
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
			go p.SetDayInc()
		case "fishing":
			go p.setFishingResult(params[1])
			go p.createHaul(params[1])
		}
	}
}

func (p *Person) getInventory() map[uuid.UUID]InventoryItem {
	return p.Inventory
}

func (p *Person) createHaul(res string) {
	var hauls []FishHaul
	err := json.Unmarshal([]byte(res), &hauls)
	if err != nil {
		fmt.Printf("При маршалинге улова в JSON у персонажа %d произошла ошибка %s", p.ID, err)
		return
	}
	for i := range hauls {
		f := getFishByID(hauls[i].ID)
		item := getItemPool().(*Item)
		item.UUID = uuid.Must(uuid.NewV1())
		item.Name = f.Name
		item.Weight = hauls[i].Weight
		item.Quality = hauls[i].Qaulity
		item.Limit = 3
		item.CreationDate = GetDate()
		item.ExpDate = item.CreationDate + item.Limit
		item.Object.IsCountable = true
		item.Object.Name = "Рыба"
		Items = append(Items, item)
		p.Inventory[item.UUID] = InventoryItem{item.UUID, GetDate()}
	}
}

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
		dM6 := float64(GetRandInt(0, 40)-20) / 100.0
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
			item.UUID,
			item.Name,
			item.Weight,
			item.Limit,
			item.Quality,
			item.CreationDate,
			item.ExpDate}
		inv = append(inv, pitem)
	}
	return
}

func GetRandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getRandMale() bool {
	return GetRandInt(0, 1) != 0
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
		return maleName[GetRandInt(0, len(maleName)-1)] + " " + maleSurname[GetRandInt(0, len(maleSurname)-1)]
	}
	return femaleName[GetRandInt(0, len(femaleName)-1)] + " " + femaleSurname[GetRandInt(0, len(femaleSurname)-1)]

}

func CreatePerson(count int) {
	Persons = make([]Person, count)
	for i := range Persons {
		isMale := getRandMale()
		Persons[i] = Person{
			uuid.Must(uuid.NewV1()),
			getRandName(isMale),
			GetRandInt(18, 28),
			isMale,
			1,
			1.0,
			make(chan string, 0),
			make(map[uuid.UUID]InventoryItem)}
	}
}

func GetPersons() []Person {
	return Persons
}
