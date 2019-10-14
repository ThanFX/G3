package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type River struct {
	Terrain
	IsBribge bool
}

type Rivers struct {
	Objects []River
	InCh    chan string `json:"-"`
}

var R Rivers

func CreateRiver(chunkId uuid.UUID, size int, isBridge bool) uuid.UUID {
	r := River{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId},
		isBridge}
	R.Objects = append(R.Objects, r)
	return r.ID
}

func GetRivers() []River {
	return R.Objects
}
