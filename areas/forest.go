package areas

import (
	uuid "github.com/satori/go.uuid"
)

type Forest struct {
	Area
}

type Forests struct {
	Objects []Forest
	InCh    chan string `json:"-"`
}

var F Forests

func CreateForest(chunkId uuid.UUID, size int) uuid.UUID {
	f := Forest{
		Area{uuid.Must(uuid.NewV4()),
			size,
			chunkId,
			[]AreaMastery{
				AreaMastery{GetMasteryByName("hunting"), 0, 0},
				AreaMastery{GetMasteryByName("food_gathering"), 0, 0}}}}
	F.Objects = append(F.Objects, f)
	return f.ID
}

func GetForests() []Forest {
	return F.Objects
}

func GetForestById(id uuid.UUID) []Forest {
	var f []Forest
	for i := range F.Objects {
		if uuid.Equal(F.Objects[i].ID, id) {
			f = append(f, F.Objects[i])
		}
	}
	return f
}
