package models

import "math/rand"

type Person struct {
	Id     int
	Name   string
	Birth  int
	IsMale bool
	Chunk  int
}

var Persons []Person

func getRandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getRandMale() bool {
	return getRandInt(0, 1) != 0
}

func getRandName(isMale bool) string {
	femaleName := []string{"Лаиммика",
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
	femaleSurname := []string{"Суивра",
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
	maleName := []string{"Беронлас",
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
	maleSurname := []string{"Кинадур",
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
		return maleName[getRandInt(0, len(maleName)-1)] + " " + maleSurname[getRandInt(0, len(maleSurname)-1)]
	} else {
		return femaleName[getRandInt(0, len(femaleName)-1)] + " " + femaleSurname[getRandInt(0, len(femaleSurname)-1)]
	}

}

func CreatePerson(count int) {
	Persons = make([]Person, count)
	for i := range Persons {
		isMale := getRandMale()
		Persons[i] = Person{
			i + 1,
			getRandName(isMale),
			getRandInt(18, 28),
			isMale,
			1}
	}
}

func GetPersons() []Person {
	return Persons
}
