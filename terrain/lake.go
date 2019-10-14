package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Lake struct {
	Terrain
}

type Lakes struct {
	Objects []Lake
	InCh    chan string `json:"-"`
}

var L Lakes

func CreateLake(chunkId uuid.UUID, size int) uuid.UUID {
	l := Lake{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId}}
	L.Objects = append(L.Objects, l)
	return l.ID
}

func GetLakes() []Lake {
	return L.Objects
}
