package models

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID     int
	Name   string
	Birth  int
	IsMale bool
	Chunk  int
	Skill  float64
	UUID   uuid.UUID   `json:"-"`
	InCh   chan string `json:"-"`
}

var Persons []Person

func (p *Person) SetDayInc() {
	lakeUUID := GetRandLakeUUID()
	LakeMessage(lakeUUID, fmt.Sprintf("fishing|%s|%s", strconv.FormatFloat(p.Skill, 'f', -1, 64), p.UUID.String()))
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
			go p.setFishingResult(params[1], params[2])
		}
	}
}

func (p *Person) setFishingResult(res, lakeId string) {
	skillUp, err := strconv.ParseFloat(res, 64)
	if err != nil {
		fmt.Printf("Ошибка парсинга улова у персонажа %s: %s", p.Name, err)
		skillUp = 0
	}

	roundedUp := math.Round(skillUp*100) / 1000
	p.Skill += roundedUp
	p.Skill = math.Floor(p.Skill*100) / 100
	NewEvent(
		fmt.Sprintf("Персонаж %s выловил %s рыбы и получил прирост навыка рыбалки на %s. Текущее значение навыка - %s",
			p.Name, res, strconv.FormatFloat(roundedUp, 'f', -1, 64), strconv.FormatFloat(p.Skill, 'f', -1, 64)))
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
		if uuid.Equal(Persons[i].UUID, id) {
			return Persons[i], nil
		}
	}
	return Persons[0], errors.New("Такой персонаж не найдено\n")
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
			i + 1,
			getRandName(isMale),
			GetRandInt(18, 28),
			isMale,
			1,
			1.0,
			uuid.Must(uuid.NewV1()),
			make(chan string, 0)}
	}
}

func GetPersons() []Person {
	return Persons
}
