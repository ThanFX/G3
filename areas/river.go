package areas

import (
	uuid "github.com/satori/go.uuid"
)

type River struct {
	Area
	IsBribge bool
}

type Rivers struct {
	Objects []River
	InCh    chan string `json:"-"`
}

var R Rivers

func CreateRiver(chunkId uuid.UUID, size int, isBridge bool) uuid.UUID {
	r := River{
		Area{uuid.Must(uuid.NewV4()),
			size,
			chunkId,
			[]AreaMastery{
				AreaMastery{GetMasteryByName("fishing"), 0, 0}}},
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
