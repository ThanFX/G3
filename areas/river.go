package areas

import (
	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type River struct {
	libs.Area
	IsBribge bool
}

type Rivers struct {
	Objects []River
	InCh    chan string `json:"-"`
}

var R Rivers

func CreateRiver(chunkId uuid.UUID, size int, isBridge bool) uuid.UUID {
	r := River{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("fishing"),
					Capacity:    0,
					MaxCapacity: 0}}},
		isBridge}
	R.Objects = append(R.Objects, r)
	return r.ID
}

func GetRivers() []River {
	return R.Objects
}

func GetRiversById(id uuid.UUID) []River {
	var r []River
	for i := range R.Objects {
		if uuid.Equal(R.Objects[i].ID, id) {
			r = append(r, R.Objects[i])
		}
	}
	return r
}
