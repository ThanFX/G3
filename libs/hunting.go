package libs

import (
	"math"
	"math/rand"
)

func GetHuntingInitSize(size int) (cap, maxCap int) {
	switch size {
	case 1:
		cap = 25
		maxCap = 25
	case 2:
		cap = 50
		maxCap = 50
	case 3:
		cap = 100
		maxCap = 100
	case 4:
		cap = 250
		maxCap = 250

	case 5:
		cap = 500
		maxCap = 500
	}
	return
}

func GetHuntingDayInc(cap, size int) int {
	_, maxCap := GetHuntingInitSize(size)
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

func getHuntWeight(mastery float64, min, max int) int {
	// Считаем массу добычи (чем выше рандом, тем выше максимальный вес)
	randWeight := GetRandInt(1, int(math.Floor(mastery)))
	// Максимальная масса добычи, которую можно заохотить с таким уровнем навыка
	fw := float64(min) + (float64(max-min) * float64(randWeight) * 0.01)
	fw = math.Floor(fw/100) * 100
	return int(fw)
}

func getHuntQuality(mastery float64) int {
	var huntQuality = 1
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
		if qual > huntQuality {
			huntQuality = qual
		}
	}
	return huntQuality
}

func getHunt(mastery float64, areaType string, areaSize int) MasteryItem {
	var huntRarity = 1
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
		if rar > huntRarity {
			huntRarity = rar
		}
	}
	return getRandMasteryItemIDFromArea("hunting", areaType, areaSize, huntRarity)
}

func CalcPersonHuntingDayHaul(areaType string, areaSize int, personSkill float64) []Haul {
	//fmt.Println(areaType, areaSize, personSkill)
	var hauls []Haul
	mastery := getMasteryFunc(personSkill)
	//NewEvent(fmt.Sprintf("Уровень навыка - %f, мастерство - %f", s, mastery))
	// 4 тика: 8 часов * 0.5 скорость
	for i := 0; i < 4; i++ {
		// шанс на улов
		f := float64(0.1) + (mastery / 100)
		if rand.Float64() < f {
			var haul Haul
			// Поймали кого-то, теперь нужно определить её параметры
			// определяем вид
			hunt := getHunt(mastery, areaType, areaSize)
			haul.ID = hunt.ID
			// Определяем массу
			haul.Weight = getHuntWeight(mastery, hunt.Min, hunt.Max)
			// Определяем качество
			haul.Qaulity = getHuntQuality(mastery)
			hauls = append(hauls, haul)
		}
	}
	return hauls
}
