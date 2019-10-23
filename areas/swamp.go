package areas

import (
	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Swamp struct {
	libs.Area
}

type Swamps struct {
	Objects []Swamp
	InCh    chan string `json:"-"`
}

var S Swamps

func CreateSwamp(chunkId uuid.UUID, size int) uuid.UUID {
	s := Swamp{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("food_gathering"),
					Capacity:    0,
					MaxCapacity: 0}}}}
	S.Objects = append(S.Objects, s)
	return s.ID
}

func GetSwamps() []Swamp {
	return S.Objects
}

func GetSwampById(id uuid.UUID) Swamp {
	var s Swamp
	for i := range S.Objects {
		if uuid.Equal(S.Objects[i].ID, id) {
			s = S.Objects[i]
			break
		}
	}
	return s
}
