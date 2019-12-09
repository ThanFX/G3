package libs

import (
	"math"
	"math/rand"
)

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

func getFish(mastery float64, areaType string, areaSize int) MasteryItem {
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
		if rar > fishRarity {
			fishRarity = rar
		}
	}
	return getRandMasteryItemIDFromArea("fishing", areaType, areaSize, fishRarity)
}

func getFishQuality(mastery float64) int {
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
		if qual > fishQuality {
			fishQuality = qual
		}
	}
	return fishQuality
}

func getFishWeight(mastery float64, min, max int) int {
	// Считаем массу рыбки (чем выше рандом, тем выше максимальный вес)
	randWeight := GetRandInt(1, int(math.Floor(mastery)))
	// Максимальная масса рыбы, которую можно выловить с таким уровнем навыка
	fw := float64(min) + (float64(max-min) * float64(randWeight) * 0.01)
	fw = math.Floor(fw/100) * 100
	return int(fw)
}

func CalcPersonFishingDayHaul(areaType string, areaSize int, personSkill float64) []Haul {
	//fmt.Println(areaType, areaSize, personSkill)
	var hauls []Haul
	mastery := getMasteryFunc(personSkill)
	//NewEvent(fmt.Sprintf("Уровень навыка - %f, мастерство - %f", s, mastery))
	// 32 тика: 8 часов * 4 скорость удочки
	for i := 0; i < 32; i++ {
		// шанс на улов
		f := float64(0.1) + (mastery / 100)
		if rand.Float64() < f {
			var haul Haul
			// Поймали рыбу, теперь нужно определить её параметры
			// определяем вид пойманой рыбы
			fish := getFish(mastery, areaType, areaSize)
			haul.ID = fish.ID
			// Определяем массу пойманной рыбы
			haul.Weight = getFishWeight(mastery, fish.Min, fish.Max)
			// Определяем качество пойманной рыбы
			haul.Qaulity = getFishQuality(mastery)
			// Добавляем рыбу в улов
			hauls = append(hauls, haul)
		}
	}
	return hauls
}
