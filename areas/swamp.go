package areas

import (
	uuid "github.com/satori/go.uuid"
)

type Swamp struct {
	Area
}

type Swamps struct {
	Objects []Swamp
	InCh    chan string `json:"-"`
}

var S Swamps

func CreateSwamp(chunkId uuid.UUID, size int) uuid.UUID {
	s := Swamp{
		Area{uuid.Must(uuid.NewV4()),
			size,
			chunkId,
			[]AreaMastery{
				AreaMastery{GetMasteryByName("food_gathering"), 0, 0}}}}
	S.Objects = append(S.Objects, s)
	return s.ID
}

func GetSwamps() []Swamp {
	return S.Objects
}

func GetSwampById(id uuid.UUID) []Swamp {
	var s []Swamp
	for i := range S.Objects {
		if uuid.Equal(S.Objects[i].ID, id) {
			s = append(s, S.Objects[i])
		}
	}
	return s
}
