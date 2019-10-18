package areas

import (
	uuid "github.com/satori/go.uuid"
)

type Hill struct {
	Area
}

type Hills struct {
	Objects []Hill
	InCh    chan string `json:"-"`
}

var H Hills

func CreateHill(chunkId uuid.UUID, size int) uuid.UUID {
	h := Hill{
		Area{uuid.Must(uuid.NewV4()),
			size,
			chunkId,
			[]AreaMastery{
				AreaMastery{GetMasteryByName("hunting"), 0, 0},
				AreaMastery{GetMasteryByName("food_gathering"), 0, 0}}}}
	H.Objects = append(H.Objects, h)
	return h.ID
}

func GetHills() []Hill {
	return H.Objects
}

func GetHillsById(id uuid.UUID) []Hill {
	var h []Hill
	for i := range H.Objects {
		if uuid.Equal(H.Objects[i].ID, id) {
			h = append(h, H.Objects[i])
		}
	}
	return h
}
