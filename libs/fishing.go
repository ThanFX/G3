package libs

import (
	"math"
	"math/rand"
)

type FishHaul struct {
	ID      int
	Weight  int
	Qaulity int
}

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

func CalcPersonFishingDayHaul(areaType string, areaSize int, personSkill float64) []FishHaul {
	//fmt.Println(areaType, areaSize, personSkill)
	var hauls []FishHaul
	mastery := getMasteryFunc(personSkill)
	//NewEvent(fmt.Sprintf("Уровень навыка - %f, мастерство - %f", s, mastery))
	// 32 тика: 8 часов * 4 скорость удочки
	for i := 0; i < 32; i++ {
		// шанс на улов
		f := float64(0.1) + (mastery / 100)
		if rand.Float64() < f {
			// Поймали рыбу, теперь нужно определить её параметры
			// Считаем максимальный улов (чем выше рандом, тем выше вес)
			randWeight := GetRandInt(1, int(math.Floor(mastery)))
			var haul FishHaul
			// Максимальная масса рыбы, которую можно выловить с таким уровнем навыка и в таком озере
			wfMax := float64(randWeight * 100)
			if wfMax < 100 {
				wfMax = 100
			} else {
				wfMax = math.Floor(wfMax/100) * 100
			}
			wf := GetRandInt(1, int(wfMax)/100) * 100
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
			haul.ID = getRandMasteryItemIDFromArea("fishing", areaType, areaSize, fishRarity)

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
			haul.Qaulity = fishQuality
			hauls = append(hauls, haul)
		}
	}
	return hauls
}
