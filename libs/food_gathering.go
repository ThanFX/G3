package libs

import (
	"math"
	"math/rand"
)

func GetFoodGatheringInitSize(size int) (cap, maxCap int) {
	switch size {
	case 1:
		cap = 25000
		maxCap = 25000
	case 2:
		cap = 50000
		maxCap = 50000
	case 3:
		cap = 100000
		maxCap = 100000
	case 4:
		cap = 250000
		maxCap = 250000
	case 5:
		cap = 500000
		maxCap = 500000
	}
	return
}

func GetFoodGatheringDayInc(cap, size int) int {
	_, maxCap := GetFoodGatheringInitSize(size)
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

func getFGWeight(mastery float64, min, max int) int {
	// Считаем массу добычи (чем выше рандом, тем выше максимальный вес)
	randWeight := GetRandInt(1, int(math.Floor(mastery)))
	// Максимальная масса добычи, которую можно заохотить с таким уровнем навыка
	w := float64(min) + (float64(max-min) * float64(randWeight) * 0.01)
	w = math.Floor(w/100) * 100
	return int(w)
}

func getFGQuality(mastery float64) int {
	var quality = 1
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
		if qual > quality {
			quality = qual
		}
	}
	return quality
}

func getFG(mastery float64, areaType string, areaSize int) MasteryItem {
	var rarity = 1
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
		if rar > rarity {
			rarity = rar
		}
	}
	return getRandMasteryItemIDFromArea("food_gathering", areaType, areaSize, rarity)
}

func CalcPersonFGDayHaul(areaType string, areaSize int, personSkill float64) []Haul {
	//fmt.Println(areaType, areaSize, personSkill)
	var hauls []Haul
	mastery := getMasteryFunc(personSkill)
	//NewEvent(fmt.Sprintf("Уровень навыка - %f, мастерство - %f", s, mastery))
	// 32 тика: 8 часов * 4 скорость
	for i := 0; i < 4; i++ {
		// шанс на улов
		f := float64(0.1) + (mastery / 100)
		if rand.Float64() < f {
			var haul Haul
			// Поймали кого-то, теперь нужно определить её параметры
			// определяем вид
			fg := getFG(mastery, areaType, areaSize)
			haul.ID = fg.ID
			// Определяем массу
			haul.Weight = getFGWeight(mastery, fg.Min, fg.Max)
			// Определяем качество
			haul.Qaulity = getFGQuality(mastery)
			hauls = append(hauls, haul)
		}
	}
	return hauls
}
