package areas

import (
	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Lake struct {
	libs.Area
}

type Lakes struct {
	Objects []Lake
	InCh    chan string `json:"-"`
}

var L Lakes

func CreateLake(chunkId uuid.UUID, size int) uuid.UUID {
	l := Lake{
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
					Mastership:  libs.GetMasteryByName("fishing"),
					Capacity:    0,
					MaxCapacity: 0}}}}
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
