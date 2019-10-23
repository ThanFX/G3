package areas

import (
	"github.com/ThanFX/G3/libs"

	uuid "github.com/satori/go.uuid"
)

type Forest struct {
	libs.Area
}

type Forests struct {
	Objects []Forest
	InCh    chan string `json:"-"`
}

var F Forests

func CreateForest(chunkId uuid.UUID, size int) uuid.UUID {
	f := Forest{
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
	F.Objects = append(F.Objects, f)
	return f.ID
}

func GetForests() []Forest {
	return F.Objects
}

func GetForestById(id uuid.UUID) Forest {
	var f Forest
	for i := range F.Objects {
		if uuid.Equal(F.Objects[i].ID, id) {
			f = F.Objects[i]
			break
		}
	}
	return f
}
