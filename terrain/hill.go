package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Hill struct {
	Terrain
}

type Hills struct {
	Objects []Hill
	InCh    chan string `json:"-"`
}

var H Hills

func CreateHill(chunkId uuid.UUID, size int) {
	h := Hill{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId}}
	H.Objects = append(H.Objects, h)
}

func GetHills() []Hill {
	return H.Objects
}
