package areas

import (
	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Meadow struct {
	libs.Area
}

type Meadows struct {
	Objects []Meadow
	InCh    chan string `json:"-"`
}

var M Meadows

func CreateMeadow(chunkId uuid.UUID, size int) uuid.UUID {
	m := Meadow{
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
	M.Objects = append(M.Objects, m)
	return m.ID
}

func GetMeadows() []Meadow {
	return M.Objects
}

func GetMeadowById(id uuid.UUID) Meadow {
	var m Meadow
	for i := range M.Objects {
		if uuid.Equal(M.Objects[i].ID, id) {
			m = M.Objects[i]
			break
		}
	}
	return m
}
