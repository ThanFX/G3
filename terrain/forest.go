package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Forest struct {
	Terrain
}

type Forests struct {
	Objects []Forest
	InCh    chan string `json:"-"`
}

var F Forests

func CreateForest(chunkId uuid.UUID, size int) uuid.UUID {
	f := Forest{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId}}
	F.Objects = append(F.Objects, f)
	return f.ID
}

func GetForests() []Forest {
	return F.Objects
}
