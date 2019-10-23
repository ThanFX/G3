package areas

import (
	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Hill struct {
	libs.Area
}

type Hills struct {
	Objects []Hill
	InCh    chan string `json:"-"`
}

var H Hills

func CreateHill(chunkId uuid.UUID, size int) uuid.UUID {
	h := Hill{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("hunting"),
					Capacity:    0,
					MaxCapacity: 0},
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("food_gathering"),
					Capacity:    0,
					MaxCapacity: 0}}}}
	H.Objects = append(H.Objects, h)
	return h.ID
}

func GetHills() []Hill {
	return H.Objects
}

func GetHillsById(id uuid.UUID) Hill {
	var h Hill
	for i := range H.Objects {
		if uuid.Equal(H.Objects[i].ID, id) {
			h = H.Objects[i]
			break
		}
	}
	return h
}
