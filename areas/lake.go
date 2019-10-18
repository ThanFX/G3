package areas

import (
	uuid "github.com/satori/go.uuid"
)

type Lake struct {
	Area
}

type Lakes struct {
	Objects []Lake
	InCh    chan string `json:"-"`
}

var L Lakes

func CreateLake(chunkId uuid.UUID, size int) uuid.UUID {
	l := Lake{
		Area{uuid.Must(uuid.NewV4()),
			size,
			chunkId,
			[]AreaMastery{
				AreaMastery{GetMasteryByName("hunting"), 0, 0},
				AreaMastery{GetMasteryByName("fishing"), 0, 0}}}}
	L.Objects = append(L.Objects, l)
	return l.ID
}

func GetLakes() []Lake {
	return L.Objects
}

func GetLakesById(id uuid.UUID) []Lake {
	var l []Lake
	for i := range L.Objects {
		if uuid.Equal(L.Objects[i].ID, id) {
			l = append(l, L.Objects[i])
		}
	}
	return l
}
